# Cisco Webex Notifications Bot

## Description
cisco-webex-bot is a CLI application written in Golang that gives the ability to send notifications to a Webex room that the bot is subscribed to via a rest api


## When to use cisco-webex-bot
- When you need to notify users in Webex. 
- Use as part of a CI/CD pipeline to give status updates.
- Use as part of an automation pipeline.


## Requirements
* go 1.8 [https://go.dev/doc/install](https://go.dev/doc/install) to run and install cisco-webex-bot
* helm if you want to rebuild the helm package
* docker if you want to build or run the container image 
* KubeConfig for the cluster you wish to deploy to.
* Git if you wish to clone cisco-webex-bot project
* Swag to update swagger document [https://github.com/swaggo/swag/cmd/swag](https://github.com/swaggo/swag/cmd/swag) 

## Project folders
Below is a description cisco-webex-bot project folders and what they contain
|   Folder        | Description  | 
|-----------|---|
| charts    | Contains the helm chart for cisco-webex-bot  |
| docs      | Contains the swagger documents |
| documents | Contains cli and api markdown files  |
| modules   | Contains cisco-webex-bot modules and code  |
| cmd       | Contains code for cisco-webex-bot CLI   |
|           |   |

## Installation and Basic usage
This will take you through the steps to install and get cisco-webex-bot up and running.
<details>
<summary>1. Install</summary>

Once you have installed golang you can run the following command to install cisco-webex-bot
```yaml
go install github.com/Mrpye/cisco-webex-bot
```
</details>

<details>

<summary>2. Start the cisco-webex-bot web server</summary>

This will run the web-server on port 8080
```bash
    cisco-webex-bot web --port 8080 --ip 0.0.0.0 
```
</details>

For more instructions on using the cisco-webex-bot CLI,
check out the CLI documentation [here](./documents/cisco-webex-bot.md)



## Build and Run cisco-webex-bot as a container
The following steps will take you through how to build and run cisco-webex-bot as a container image.

<details>
<summary>1. Clone the repository</summary>

This will clone the cisco-webex-bot project from github
```bash
# clone the project
git clone https://github.com/Mrpye/cisco-webex-bot.git

# Change into the directory
cd cisco-webex-bot
```
</details>

<details>
<summary>2. Build the cisco-webex-bot container image</summary>

This will build the container image you will need docker installed to build.
```
sudo docker build . -t  cisco-webex-bot:v1.0.0 -f Dockerfile
```
</details>

<details>
<summary>3. Run the container</summary>

This will run the cisco-webex-bot container and expose the API endpoint on port 8080.
```
sudo docker run -d -p 8080:8080 --name=cisco-webex-bot  --env=WEB_IP=0.0.0.0 -t cisco-webex-bot:1.0.0
```
</details>

---

## Environment  variables
as well as parameters that you can pass to cisco-webex-bot via the CLI you can also configure cisco-webex-bot using environmental variables.
- WEB_IP (set the listening ip address 0.0.0.0 allow from everywhere)
- WEB_PORT (set the port to listen on)


## cisco-webex-bot Helm chart
This guide will show you how to build the helm chart package for cisco-webex-bot, you will need to have helm installed to build the package.

<details>
<summary>1. Build the helm chart package for cisco-webex-bot</summary>

```bash
# change into the chart directory
cd charts
# Package the cisco-webex-bot chart
helm package cisco-webex-bot

```

the helm chart package will be saved under the charts folder cisco-webex-bot-0.1.0.tgz

</details>

<details>
<summary>2. Configure cisco-webex-bot chart</summary>
below are main setting you may want to modify

```yaml
image:
  repository: cisco-webex-bot
  pullPolicy: Always
  # Overrides the image tag whose default is the chart appVersion.
  tag: "v1.0.0"

#Set the container env values
WebServer:
  listenOn: "0.0.0.0"
  port: "8080"

```

</details>


<details>
<summary>3. How to configure and deploy using cisco-webex-bot</summary>
Based on the configuration setting from previous step this is what a payload would look like when installing using cisco-webex-bot

```
# Change to thew charts directory
cd charts

#run the chart
helm install cisco-webex-bot --values values.yaml
```
</details>

---
## Create a Bot token
To be able to use the cisco-webex-bot you will need to get a token.

- you can create the token here
[https://developer.webex.com/docs/bots](https://developer.webex.com/docs/bots)
- click on the Create a Bot.
- you will have to login with a Cisco Account.
- Give your Bot a name.
- Give your bot a username.
- Upload an icon.
- Description of what your bot does.
- Then click Add Bot
- Save the token somewhere safe you will need this to make calls.



## Using the API
The following guide will show you how to make call to the various API endpoints with examples. 

- you can also refer to the API document [here](./documents/api.md) for more details. 
- Also there is a swagger interface you can access when cisco-webex-bot web-server is running. 

```
http://localhost:8080/docs/index.html
```

## Example API calls

<details>
<summary>Send Message Text format</summary>

``` bash
curl --location --request POST '172.16.10.237:9040/send_message' \
--header 'Authorization: Bearer <webex token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message":"hello world"
}'
```

</details>

<details>
<summary>Send Message as Markdown</summary>

``` bash
curl --location --request POST '172.16.10.237:9040/send_message_md' \
--header 'Authorization: Bearer <webex token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message":"# hello world \n **This is a test**"
}'
```

</details>

<details>
<summary>Test the Web Server is Alive</summary>

```bash
curl --location --request GET 'localhost:8080/'
```
Return OK

</details>

---

## Update the swagger document
The code below shows you how to update the swagger API documents.

If you need more helm on using these tools please refer to the links below
- gin-swagge [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)
- swag [https://github.com/swaggo/swag](https://github.com/swaggo/swag)

<details>
<summary>1. Install swag</summary>

```bash
#Install swag
go install github.com/swaggo/swag/cmd/swag
```
</details>

<details>
<summary>2. Update APi document</summary>

```bash
#update the API document
swag init
```
</details>

<details>
<summary>3. Update the api.md</summary>

```bash
swagger generate markdown -f .\docs\swagger.json --output .\documents\api.md 
```
</details>

---

## To Do
- Nothing at the moment

--- 

## Main 3rd party Libraries

- [https://github.com/helm/helm](https://github.com/helm/helm)
- [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger) 
- [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)


## license
cisco-webex-bot is Apache 2.0 licensed.