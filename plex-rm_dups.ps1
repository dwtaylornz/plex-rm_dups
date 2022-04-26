#Note : this script differs slightly from parent...
#       check score_video fuction for scoring

$Server='ip:port'
$PlexToken = 'tokengoeshere'
$libraries = Invoke-RestMethod -Uri "http://$server/library/sections/all?X-Plex-Token=$PlexToken"
$VideoLibraries = $libraries.MediaContainer.Directory 

function process_episode([object]$episode) {

    Write-Host -ForegroundColor Yellow "$($episode.grandparentTitle) $($episode.title) $($episode.Key)"          
    $videos = Invoke-RestMethod -Uri "http://$server$($episode.key)?X-Plex-Token=$PlexToken"
    $loser = score_videos($videos)

    Write-Host "  Loser id:" $loser.Key -NoNewline
    $url = "http://$Server$($Episode.Key)/media/$($Loser.Key)?X-Plex-Token=$PlexToken" 
    Invoke-RestMethod -Method Delete -Uri $url   # comment me out if you want to test first!
    Write-Host "  ...deleted"
}

function score_videos ([object]$videos) {

    $scores = @{}
    
    foreach ($media in $videos.MediaContainer.Video.Media) {

        $scores[$media.id] = 0 

        if (!$first_video_size) { $first_video_size = $media.Part.size }

        #video codec
        if ($media.videoCodec -eq "hevc") { $scores[$media.id] += 1000 }
        if ($media.videoCodec -eq "h264") { $scores[$media.id] += 500 }

        #audio codec
        if ($media.audioCodec -eq "aac") { $scores[$media.id] += 10 }
    
        #resolution
        if ($media.videoResolution -eq "1080") { $scores[$media.id] += 1000 }
        if ($media.videoResolution -eq "720") { $scores[$media.id] += 500 }
        
        #size
        if ($media.Part.size -lt $first_video_size) { $scores[$media.id] += 100 }
        if ($media.Part.size -gt $first_video_size) { $scores[$media.id] -= 100 }

        Write-Host "  id:" $media.id "video_codec:" $media.videoCodec "audio_codec:" $media.audioCodec "resoultion:" $media.videoResolution "size:" $media.Part.size "score:" $scores[$media.id]
    }

    $loser = $scores.GetEnumerator() | Sort-Object Value -Descending | Select-Object -Last 1
    return $loser
}

foreach ($library in $VideoLibraries) {

    Write-Host -ForegroundColor Yellow $library.title 
    
    if ($($library.type) -eq "movie") {    
        $episodes = Invoke-RestMethod -Uri "http://$server/library/sections/$($library.key)/all?duplicate=1&X-Plex-Token=$PlexToken"
        foreach ($episode in $episodes.MediaContainer.Video) {       
            process_episode $episode           
        }
    }

    if ($($library.type) -eq "show") {
        $episodes = Invoke-RestMethod -Uri "http://$server/library/sections/$($library.key)/search?type=4&duplicate=1&X-Plex-Token=$PlexToken"
        foreach ($episode in $episodes.MediaContainer.Video) {
            process_episode $episode
        } 
    }
}
