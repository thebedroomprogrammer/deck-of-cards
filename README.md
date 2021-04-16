# Deck of Cards

## Description
A simple backend server that extends a utility via APIs to manage a deck of cards and provide functionality to perform basic operations on the deck like opening a deck, creating a deck and drawing cards from a deck.

## API

### 1) Create a deck
It would create the standard 52-card deck of French playing cards, It includes
all thirteen ranks in each of the four suits: clubs (♣), diamonds (♦), hearts (♥)
and spades (♠).

#### Request
`GET /create` - Creates an unshuffled deck with 52 cards

`GET /create?cards=AS,KH&shuffle=true` - Creates an shuffled deck with 2 cards

#### Response
```
{
	"deck_id": "3b91b2fd-173c-4b01-8fb6-1fe1c2d6f5ec",
	"shuffled": false,
	"remaining": 52
}
```


### 2) Draw Cards from a deck
It would draw a card(s) of a given Deck. If the deck was not passed over or
invalid it should return an error. A count parameter needs to be provided to
define how many cards to draw from the deck.

#### Request
`GET /draw?deck_id=3b91b2fd-173c-4b01-8fb6-1fe1c2d6f5ec&count=2` - Draws two cards from the deck with the provided deck Id.

#### Response
```
{
    "cards": [
        {
            "value": "6",
            "suit": "SPADES",
            "code": "6S"
        },
        {
            "value": "7",
            "suit": "SPADES",
            "code": "7S"
        }
    ]
}
```

### 3) Open a deck
It would return a given deck by its UUID. If the deck was not passed over or is
invalid it should return an error. This method will "open the deck", meaning that
it will list all cards by the order it was created.

#### Request
`GET /open?deck_id=3b91b2fd-173c-4b01-8fb6-1fe1c2d6f5ec` - List all cards present in the deck in a detailed manner

#### Response
```
{
    "deck_id": "3b91b2fd-173c-4b01-8fb6-1fe1c2d6f5ec",
    "shuffled": false,
    "remaining": 2,
    "cards": [
        {
            "value": "QUEEN",
            "suit": "HEARTS",
            "code": "QH"
        },
        {
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        }
    ]
}
```

## Testing
The test output pattern follows Given-when-then or Given-when-should pattern. More about it in here
https://www.codeguru.com/cpp/sample_chapter/go-in-action-basic-unit-test.html

`make test` - tests both the utilities, handlers and store

`make test-utils` - tests utilities

`make test-store` - tests in memory data store

`make test-handlers` - tests all the handlers


## Running Locally
In the root directory run the command mentioned below to start the server
`make run`

This will start the server at http://127.0.0.1/8080

## Making build
In the root directory run the command mentioned below to start the server
`make build`

This will output a build file in the root folder.