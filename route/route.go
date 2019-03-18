package route

import (
    "github.com/gorilla/mux"
    messaging "github.com/heaptracetechnology/microservice-telegram/messaging"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "GetBotDetails",
        "GET",
        "/getbot",
        messaging.GetBotDetails,
    },
    Route{
        "SendGroupMessage",
        "POST",
        "/sendgroupmessage",
        messaging.SendGroupMessage,
    },
    Route{
        "SendChannelMessage",
        "POST",
        "/sendchannelmessage",
        messaging.SendChannelMessage,
    },
    Route{
        "GetChat",
        "POST",
        "/getchat",
        messaging.GetChat,
    },
    Route{
        "LeaveChat",
        "POST",
        "/leavechat",
        messaging.LeaveChat,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
