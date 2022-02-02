package restchat

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetOnlineUserIds(t *testing.T) {
	sessionmemrep := new(UserSessionsMemRepo)
	sessionmemrep.Sessions = append(sessionmemrep.Sessions, SessionModel{ID: 1, Username: "Vasya", Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	sessionmemrep.Sessions = append(sessionmemrep.Sessions, SessionModel{ID: 2, Username: "Kolya", Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})
	userid := []int{1, 2}
	request, err := sessionmemrep.GetOnlineUserIds()
	if reflect.DeepEqual(request, userid) != false {
		t.Errorf("Не верный список пользователей")
	}

	fmt.Printf("Все хорошо! массив готов %v\n", err)
}
func TestDeleteSessionByIndex(t *testing.T) {

	sessionmemrep := new(UserSessionsMemRepo)
	sessionmemrep.Sessions = append(sessionmemrep.Sessions, SessionModel{ID: 1, Username: "Vasya", Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	sessionmemrep.Sessions = append(sessionmemrep.Sessions, SessionModel{ID: 2, Username: "Kolya", Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})

	outusersession := new(UserSessionsMemRepo)
	outusersession.Sessions = append(outusersession.Sessions, SessionModel{ID: 1, Username: "Vasya", Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	request := deleteSessionByIndex(sessionmemrep.Sessions, 1)

	if reflect.DeepEqual(request, outusersession.Sessions) != true {
		t.Errorf("\nНе верно удалилась запись в массиве\n%v\n%v", request, outusersession.Sessions)
	}
	fmt.Printf("По указанному id удалили массив %v\n", outusersession.Sessions)
}
func TestReceivelastIDSessionEmpty(t *testing.T) {
	usmr := new(UserSessionsMemRepo)
	if getLastSessionId(usmr) != 0 {
		t.Errorf("Массив сообщений не пустой")
	}
	fmt.Printf("Все хорошо!! массив пустой %v\n", getLastSessionId(usmr))
}
func TestReceivelastIDSession(t *testing.T) {

	usmr := new(UserSessionsMemRepo)
	usmr.Sessions = append(usmr.Sessions, SessionModel{ID: 4, Username: "1", Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	usmr.Sessions = append(usmr.Sessions, SessionModel{ID: 2, Username: "3", Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})
	if getLastSessionId(usmr) != 4 {
		t.Errorf("Не верное выводиться id последней сессии")
	}
	fmt.Printf("Все хорошо!! Идентификатор последней сессии: %v\n", getLastSessionId(usmr))
}

func TestDeleteSessionEmpty(t *testing.T) {
	usmr := new(UserSessionsMemRepo)
	api_token := ""
	request, err := usmr.DeleteSession(api_token)
	if request.Sessions != nil || len(request.Sessions) != 0 {
		t.Errorf("Не удалось удалить сесcию, пустой токен")
	}

	fmt.Printf("%v\n", err)
}
