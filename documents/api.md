


# cisco-webex-bot
cisco-webex-bot is a CLI application written in Golang that gives the ability to send notifications to a Webex room that the bot is subscribed to via a rest api
  

## Informations

### Version

1.0

### License

[Apache 2.0 licensed](https://github.com/Mrpye/cisco-webex-bot/blob/main/LICENSE)

### Contact

  https://github.com/Mrpye/cisco-webex-bot

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | / | [check api endpoint](#check-api-endpoint) | Check API Endpoint |
| POST | /send_message_md | [post send message markdown](#post-send-message-markdown) | Send  a message to Webex markdown format |
| POST | /send_message | [post send message text](#post-send-message-text) | Send  a message to Webex text format |
  


## Paths

### <span id="check-api-endpoint"></span> Check API Endpoint (*check-api-endpoint*)

```
GET /
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#check-api-endpoint-200) | OK | ok |  | [schema](#check-api-endpoint-200-schema) |

#### Responses


##### <span id="check-api-endpoint-200"></span> 200 - ok
Status: OK

###### <span id="check-api-endpoint-200-schema"></span> Schema
   
  



### <span id="post-send-message-markdown"></span> Send  a message to Webex markdown format (*post-send-message-markdown*)

```
POST /send_message_md
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| request | `body` | [BodyTypesSendMessage](#body-types-send-message) | `models.BodyTypesSendMessage` | | ✓ | | query params |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-send-message-markdown-200) | OK | message sent |  | [schema](#post-send-message-markdown-200-schema) |
| [404](#post-send-message-markdown-404) | Not Found | error |  | [schema](#post-send-message-markdown-404-schema) |

#### Responses


##### <span id="post-send-message-markdown-200"></span> 200 - message sent
Status: OK

###### <span id="post-send-message-markdown-200-schema"></span> Schema
   
  



##### <span id="post-send-message-markdown-404"></span> 404 - error
Status: Not Found

###### <span id="post-send-message-markdown-404-schema"></span> Schema
   
  



### <span id="post-send-message-text"></span> Send  a message to Webex text format (*post-send-message-text*)

```
POST /send_message
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| request | `body` | [BodyTypesSendMessage](#body-types-send-message) | `models.BodyTypesSendMessage` | | ✓ | | query params |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-send-message-text-200) | OK | message sent |  | [schema](#post-send-message-text-200-schema) |
| [404](#post-send-message-text-404) | Not Found | error |  | [schema](#post-send-message-text-404-schema) |

#### Responses


##### <span id="post-send-message-text-200"></span> 200 - message sent
Status: OK

###### <span id="post-send-message-text-200-schema"></span> Schema
   
  



##### <span id="post-send-message-text-404"></span> 404 - error
Status: Not Found

###### <span id="post-send-message-text-404-schema"></span> Schema
   
  



## Models

### <span id="body-types-send-message"></span> body_types.SendMessage


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| message | string| `string` |  | |  |  |


