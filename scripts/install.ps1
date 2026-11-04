#Requires -Version 7
<#
.SYNOPSIS
    Instala o claude-tools.exe na pasta de programas e o põe no PATH do usuário.

.DESCRIPTION
    Copia bin\claude-tools.exe (gerado por 'go tool task build') para
    <Root>\claude-tools-v<VERSION>\, recria a junction <Root>\current apontando para
    essa pasta e acrescenta <Root>\current ao PATH do usuário, se ainda não estiver lá.
    Rodar de novo dá o mesmo resultado.

.PARAMETER Root
    Pasta geral do programa (ex.: C:\dev\programs\claude-tools). Obrigatório: não há
    pasta padrão.

.PARAMETER SkipPath
    Não mexe no PATH do usuário.
#>
[CmdletBinding()]
param(
    [string]$Root,
    [switch]$SkipPath
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

# Avisa as janelas abertas (o Explorer, em especial) que o ambiente mudou, para que
# os terminais abertos a partir dele já venham com o PATH novo.
function Send-EnvironmentChange {
    if (-not ('ClaudeTools.NativeMethods' -as [type])) {
        Add-Type -Namespace ClaudeTools -Name NativeMethods -MemberDefinition @'
[DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Unicode)]
public static extern IntPtr SendMessageTimeout(IntPtr hWnd, uint Msg, UIntPtr wParam, string lParam, uint fuFlags, uint uTimeout, out UIntPtr lpdwResult);
'@
    }
    $HWND_BROADCAST = [IntPtr]0xffff
    $WM_SETTINGCHANGE = 0x1A
    $SMTO_ABORTIFHUNG = 2
    $result = [UIntPtr]::Zero
    [ClaudeTools.NativeMethods]::SendMessageTimeout($HWND_BROADCAST, $WM_SETTINGCHANGE, [UIntPtr]::Zero, 'Environment', $SMTO_ABORTIFHUNG, 5000, [ref]$result) | Out-Null
}

# Acrescenta $Entry ao PATH do usuário, se ainda não estiver lá. Lê e grava o
# registro direto, sem expandir as variáveis (%USERPROFILE%...) e mantendo o tipo do
# valor: [Environment]::SetEnvironmentVariable gravaria um REG_SZ já expandido.
# Devolve $true quando acrescentou.
function Add-UserPath([string]$Entry) {
    $key = [Microsoft.Win32.Registry]::CurrentUser.OpenSubKey('Environment', $true)
    if (-not $key) { throw 'Não consegui abrir HKCU\Environment.' }
    try {
        $raw = [string]$key.GetValue('Path', '', [Microsoft.Win32.RegistryValueOptions]::DoNotExpandEnvironmentNames)
        $wanted = [Environment]::ExpandEnvironmentVariables($Entry).Trim().TrimEnd('\')
        foreach ($part in $raw -split ';') {
            $expanded = [Environment]::ExpandEnvironmentVariables($part).Trim().TrimEnd('\')
            if ($expanded -ieq $wanted) { return $false }
        }
        $kind = if ($key.GetValueNames() -contains 'Path') { $key.GetValueKind('Path') } else { [Microsoft.Win32.RegistryValueKind]::ExpandString }
        $new = if ($raw.Trim() -eq '') { $Entry } else { $raw.TrimEnd(';') + ';' + $Entry }
        $key.SetValue('Path', $new, $kind)
    } finally {
        $key.Dispose()
    }
    Send-EnvironmentChange
    return $true
}

try {
    # Sem Mandatory de propósito: com ele o pwsh pediria a pasta no console em vez de falhar.
    if ([string]::IsNullOrWhiteSpace($Root)) {
        throw 'informe a pasta de instalação: -Root <pasta> (pelo Task: go tool task install DIR=<pasta>).'
    }
    $repo = Split-Path -Parent $PSScriptRoot
    $version = (Get-Content -LiteralPath (Join-Path $repo 'VERSION') -Raw).Trim()
    if ($version -notmatch '^\d+\.\d+\.\d+$') { throw "VERSION inválida: '$version'." }
    $exe = Join-Path $repo 'bin\claude-tools.exe'
    if (-not (Test-Path -LiteralPath $exe -PathType Leaf)) { throw "Não achei $exe. Rode 'go tool task build' antes." }

    $Root = [IO.Path]::GetFullPath($Root)
    $versionDir = Join-Path $Root "claude-tools-v$version"
    $current = Join-Path $Root 'current'

    New-Item -ItemType Directory -Force -Path $versionDir | Out-Null
    Copy-Item -LiteralPath $exe -Destination (Join-Path $versionDir 'claude-tools.exe') -Force

    $existing = Get-Item -LiteralPath $current -Force -ErrorAction SilentlyContinue
    if ($existing) {
        if (-not ($existing.Attributes -band [IO.FileAttributes]::ReparsePoint)) {
            throw "$current existe e não é uma junction; não vou apagá-lo."
        }
        # Delete não recursivo remove só o link: a pasta para onde ele aponta fica.
        [IO.Directory]::Delete($current, $false)
    }
    New-Item -ItemType Junction -Path $current -Target $versionDir | Out-Null

    "claude-tools $version instalado em $versionDir"
    "Junction: $current -> $versionDir"
    if ($SkipPath) {
        'PATH do usuário: não mexi (-SkipPath).'
    } elseif (Add-UserPath $current) {
        "PATH do usuário: acrescentei $current. Abra um terminal novo para usar o 'claude-tools'."
    } else {
        "PATH do usuário: $current já estava lá."
    }
} catch {
    [Console]::Error.WriteLine("Erro: $($_.Exception.Message)")
    exit 1
}
