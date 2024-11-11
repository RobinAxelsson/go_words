#!/usr/bin/env pwsh

go build .
.\go_words.exe
Write-Output ---

.\go_words.exe list
Write-Output ---

.\go_words.exe add
Write-Output ---

.\go_words.exe rnum
Write-Output ---

.\go_words.exe remove
Write-Output ---

.\go_words.exe add test
Write-Output ---

.\go_words.exe add horse stable
Write-Output ---

.\go_words.exe add dog
Write-Output ---

.\go_words.exe remove horse stable
Write-Output ---

.\go_words.exe rnum 2
Write-Output ---

.\go_words.exe rnum 1
Write-Output ---

.\go_words.exe rnum 1
Write-Output ---

.\go_words.exe rnum 99
Write-Output ---

.\go_words.exe list
Write-Output ---

.\go_words.exe remove not here
Write-Output ---

#clean up
Remove-Item wordlist.txt
Remove-Item go_words.exe