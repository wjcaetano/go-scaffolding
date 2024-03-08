# Goland

Goland won't need any further plugins to execute the APP, if you manually set env vars. For better quality-of-life you can install EnvFile plugin to support `.env` files when running the APP.

## EnvFile Plugin

#### Install & configure
1. Open Goland and go to `Preferences > Plugins` and search for `EnvFile`.
2. Install the plugin and restart Goland.

#### Running the APP
1. Go to `Run > Edit Configurations...` and select `+` to add a new configuration.
2. Select `Go Build` and name it `Build Name of the APP`.
3. In `Environment` variables, add the following:
```
## $PWD is the project root folder
APP_PATH=$PWD;configFileName=configFileName=resources/config/local.properties
```
4. Click on `EnvFile` tab and add select `+` to add new envfile
5. Choose `variables.env` on project root folder
6. Click on `Apply` and `OK`.
7. Run the APP by clicking the green play button in the top right corner.
8. You should now be able to access the APP at `http://localhost:8080`.

## Running the Integration Tests in IDE
1. Execute the last configuration created in the previous section.
2. Go to `Run > Edit Configurations...` and select `+` to add a new configuration.
3. In `Go tool arguments` add the following:

```
-tags=integration -p=1
```

### Run by the make command
If you don't want to debug the app, you can run it by the make command. To do this, you just to run the command `make run_local` in the root of the project. This command will run the app in `localhost:8080`.
