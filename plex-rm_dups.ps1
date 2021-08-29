$Server = 'server:port'
$PlexToken = 'yourtokenhere'
$libraries = Invoke-RestMethod -Uri "http://$server/library/sections/all?X-Plex-Token=$PlexToken"
$VideoLibraries = $libraries.MediaContainer.Directory 

Clear-Host

function process_episode([object]$episode) {

    Write-Host -ForegroundColor Yellow "$($episode.grandparentTitle) $($episode.title) $($episode.Key)"          
        
    $videos = Invoke-RestMethod -Uri "http://$server$($episode.key)?X-Plex-Token=$PlexToken"

    $videos.MediaContainer.Video.Media | Sort-Object -Property videoResolution, bitrate | Format-Table
    $loser = $videos.MediaContainer.Video.Media | Sort-Object -Property videoResolution, bitrate | Select-Object -Last 1
    
    Write-Output "Loser - " 
    $loser | Format-Table
    $url = "http://$Server$($Episode.Key)/media/$($Loser.id)?X-Plex-Token=$PlexToken" 
    Invoke-RestMethod -Method Delete -Uri $url
}

foreach ($library in $VideoLibraries) {

    Write-Host -ForegroundColor Yellow $library.title 
    
    if ($($library.type) -eq "movie") {    
        
        $episodes = Invoke-RestMethod -Uri "http://$server/library/sections/$($library.key)/all?duplicate=1&X-Plex-Token=$PlexToken"

        foreach ($episode in $episodes.MediaContainer.Video) {
        
            #  Write-Host "Movie Time"
            
            process_episode $episode           
        }
    }

    if ($($library.type) -eq "show") {
        $episodes = Invoke-RestMethod -Uri "http://$server/library/sections/$($library.key)/search?type=4&duplicate=1&X-Plex-Token=$PlexToken"
        foreach ($episode in $episodes.MediaContainer.Video) {
        
            # Write-Host "Show Time"

            process_episode $episode
        } 
    }
}