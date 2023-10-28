package handler

import (
	"log"
	"net/http"
	"onden-backend/api/models"
	"onden-backend/db"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var clients = make(map[string]map[*models.ClientModel]bool);

func WsConnectHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token);
	claims := user.Claims.(jwt.MapClaims);
	log.Println(claims["user_id"]);
	var userObj *models.User;
	db.DB.Where("id = ?", claims["user_id"]).First(&userObj);
	if userObj.Name == "" {
		return c.JSON(http.StatusUnauthorized, "unauthorized");
	}
	username := userObj.Name;
	roomname := c.QueryParam("room");

	if username == "" || roomname == "" {
		return c.JSON(400, "userName or roomName is empty")
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true;
		},
	}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil);
	if err != nil {
		return c.JSON(500, err.Error())
	}
	client := &models.ClientModel{Conn: conn,UserName:  username,RoomName:  roomname };

	if clients[roomname] == nil {
		clients[roomname] = make(map[*models.ClientModel]bool);
	}
	clients[roomname][client] = true;

	userList := getCurrentUsersList(roomname);

	if err := client.Conn.WriteJSON(userList); err != nil {
		return c.JSON(500, err.Error())
	}
	for {
		messageType, _, err := client.Conn.ReadMessage();
		if err != nil {
			log.Println("ReadMessage error:", err);
			break;
		}
		if err := conn.WriteMessage(messageType, []byte(username+"joined.")); err != nil {
			log.Println("WriteMessage error:", err);
			break;
		}
	}

	delete(clients[roomname], client);
	conn.Close();
	return nil;
}

func WsDisconnectHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token);
	claims := user.Claims.(jwt.MapClaims);
	log.Println(claims["user_id"]);
	var userObj *models.User;
	db.DB.Where("id = ?", claims["user_id"]).First(&userObj);
	if userObj.Name == "" {
		return c.JSON(http.StatusUnauthorized, "unauthorized");
	}
	username := userObj.Name;
	roomname := c.QueryParam("room");
	if username == "" || roomname == "" {
		return c.JSON(400, "userName or roomName is empty")
	}

	if clients[roomname] != nil {
		for client := range clients[roomname] {
			if client.UserName == username {
				delete(clients[roomname], client);
				client.Conn.Close();
				break;
			}
		}
	}

	return c.JSON(http.StatusOK, getCurrentUsersList(roomname));
}

func getCurrentUsersList(roomName string) []string {
	if clients[roomName] == nil {
		return []string{}
	}

	usersList := make([]string, 0, len(clients[roomName]))
	for client := range clients[roomName] {
		usersList = append(usersList, client.UserName)
	}
	return usersList
}