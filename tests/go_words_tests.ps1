#!/usr/bin/env pwsh

cd src
go build .
.\wordlist.exe
Write-Output ---

.\wordlist.exe list
Write-Output ---

.\wordlist.exe add
Write-Output ---

.\wordlist.exe rnum
Write-Output ---

.\wordlist.exe remove
Write-Output ---

.\wordlist.exe add test
Write-Output ---

.\wordlist.exe add horse stable
Write-Output ---

.\wordlist.exe add dog
Write-Output ---

.\wordlist.exe remove horse stable
Write-Output ---

.\wordlist.exe rnum 2
Write-Output ---

.\wordlist.exe rnum 1
Write-Output ---

.\wordlist.exe rnum 1
Write-Output ---

.\wordlist.exe rnum 99
Write-Output ---

.\wordlist.exe list
Write-Output ---

.\wordlist.exe remove not here
Write-Output ---

#clean up
Remove-Item wordlist.txt
Remove-Item wordlist.exe
cd ..
