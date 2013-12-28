Raspimusic
==========

Music player for raspberrypi - Created in project-x hackathon 

--- 

##### Install 

1. Install mpd in RPi

    ```
    apt-get update
    apt-get updgrade
    apt-get install mpd
    ```

1. Configure MPD 

	Edit configuration file ```sudo vi /etc/mpd.conf``` and update the following
	```
		 music_directory         "/path/to/music"
		 bind_to_address         "any"
	```
    And restart mpd service ```sudo service mpd restart```
    
1.  Copy the raspimusic app into RPi and run , you can change the port using ```--port``` ( Default port is 8080)
2.  Now you can access the RaspiMusic from any device using http://RPi_ip_address:port

