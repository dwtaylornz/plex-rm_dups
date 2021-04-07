$libraries = Invoke-RestMethod -Uri "http://$Server/library/sections/all?X-Plex-Token=$PlexToken"
$VideoLibraries = $libraries.MediaContainer.Directory 

foreach($library in $VideoLibraries){

    Write-Host $library.title
    $episodes = Invoke-RestMethod -Uri "http://$Server/library/sections/$($library.key)/all?duplicate=1&X-Plex-Token=$PlexToken"
    foreach($episode in $episodes.MediaContainer.Video){
        
        Write-Host "$($episode.GrandparentTitle) $($episode.ParentTitle) $($episode.title) $($episode.size)"
        
        $MediaFilesByQuality = $episode.Media | Select-Object  @{L='width';E={[int]$_.width}}, @{L='Bitrate';E={[int]$_.bitrate}}, videoCodec, id | Sort-Object -Property @{Expression = "width"; Descending = $True}, @{Expression = "bitrate"; Descending = $False}
        $Winner = $MediaFilesByQuality[0]   
        Write-Host "`tWinner is $($Winner.width) using $($Winner.videoCodec) at $($Winner.bitrate)bps"
        
        $Losers = $MediaFilesByQuality[1..$MediaFilesByQuality.Length]
        
        Foreach($Loser in $Losers){

            Write-Host "`t`tDeleting loser: $($Loser.width) using $($Loser.videoCodec) at $($Loser.bitrate)bps"
            
            $url = "http://$Server$($Episode.Key)/media/$($Loser.id)?X-Plex-Token=$PlexToken"
            
            # Write-Host $url        
            Invoke-RestMethod $url -Method Delete

            }
    }
}
