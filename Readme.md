# Dayz Server Tools

Easy management of Dayz's servers

## How to use

1. Put `server-tools.exe` in an empty folder
2. Create the `config.yaml` file in this folder
3. Fill in the `config.yaml` file with these items

```yaml
server_name: "ServerTest"
port: 2302
mission: "dayzOffline.chernarusplus"
client_mods: ["@CF"]
server_mods: []
```

4. Start `server-tools.exe` once, then make sure you have the folders `mpmissions`, `profiles`, and `serverCfgs` in your directory.
5. Note the `server_name` in your `config.yaml`, and then create a new .cfg file with the same name as it. For example, in my code above, my `server_name` is `ServerTest`, so I will create a new `ServerTest.cfg` file in the `serverCfgs` folder
6. Fill out this cfg file, official reference:

```cfg
hostname = "EXAMPLE NAME";  // Server name
password = "";              // Password to connect to the server
passwordAdmin = "";         // Password to become a server admin

description = "";			// Description of the server. Gets displayed to users in client server browser.

enableWhitelist = 0;        // Enable/disable whitelist (value 0-1)

maxPlayers = 60;            // Maximum amount of players

verifySignatures = 2;       // Verifies .pbos against .bisign files. (only 2 is supported)
forceSameBuild = 1;         // When enabled, the server will allow the connection only to clients with same the .exe revision as the server (value 0-1)

disableVoN = 0;             // Enable/disable voice over network (value 0-1)
vonCodecQuality = 20;       // Voice over network codec quality, the higher the better (values 0-30)

shardId = "123abc";			// Six alphanumeric characters for Private server

disable3rdPerson=0;         // Toggles the 3rd person view for players (value 0-1)
disableCrosshair=0;         // Toggles the cross-hair (value 0-1)

disablePersonalLight = 1;   // Disables personal light for all clients connected to server
lightingConfig = 0;         // 0 for brighter night setup, 1 for darker night setup

serverTime="SystemTime";    // Initial in-game time of the server. "SystemTime" means the local time of the machine. Another possibility is to set the time to some value in "YYYY/MM/DD/HH/MM" format, f.e. "2015/4/8/17/23" .
serverTimeAcceleration=12;  // Accelerated Time (value 0-24)// This is a time multiplier for in-game time. In this case, the time would move 24 times faster than normal, so an entire day would pass in one hour.
serverNightTimeAcceleration=1;  // Accelerated Nigh Time - The numerical value being a multiplier (0.1-64) and also multiplied by serverTimeAcceleration value. Thus, in case it is set to 4 and serverTimeAcceleration is set to 2, night time would move 8 times faster than normal. An entire night would pass in 3 hours.
serverTimePersistent=0;     // Persistent Time (value 0-1)// The actual server time is saved to storage, so when active, the next server start will use the saved time value.

guaranteedUpdates=1;        // Communication protocol used with game server (use only number 1)

loginQueueConcurrentPlayers=5;  // The number of players concurrently processed during the login process. Should prevent massive performance drop during connection when a lot of people are connecting at the same time.
loginQueueMaxPlayers=500;       // The maximum number of players that can wait in login queue

instanceId = 1;             // DayZ server instance id, to identify the number of instances per box and their storage folders with persistence files

storageAutoFix = 1;         // Checks if the persistence files are corrupted and replaces corrupted ones with empty ones (value 0-1)


class Missions
{
    class DayZ
    {
        template="dayzOffline.chernarusplus"; // Mission to load on server startup. <MissionName>.<TerrainName>
					      // Vanilla mission: dayzOffline.chernarusplus
					      // DLC mission: dayzOffline.enoch
    };
};
```

7. Put your map mission file in the mpmissions folder and make sure its name is the same as the name you have in the configuration file.

**！Please note that the map configured in the cfg will not take effect 👇（This code below will not work）**

```
class Missions
{
    class DayZ
    {
        template="dayzOffline.chernarusplus"; // Mission to load on server startup. <MissionName>.<TerrainName>
					      // Vanilla mission: dayzOffline.chernarusplus
					      // DLC mission: dayzOffline.enoch
    };
};
```
8. Run the exe file and enjoy.