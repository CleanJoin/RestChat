package restchat

import (
	"reflect"
	"testing"
)

//test написать
func TestGetUserId(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))

	var inter ISessionStorage = NewSessionStorageMemory(new(TokenGeneratorUUID))

	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 2, AuthToken: "a396776f58b942fb9b10ebc798ab6303"})

	request, err := inter.GetUserId(sessionStorage.Sessions[0].AuthToken)
	if request != 2 {
		t.Errorf("Пользователь не найден по токену  %v", sessionStorage.Sessions[0].AuthToken)
	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetUserIdBadAuthToken(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 2, AuthToken: "a396776f58b942fb9b10ebc798ab6303"})

	request, err := sessionStorage.GetUserId("a396776f58b942fb9b10ebc798ab630")

	if err == nil {
		t.Errorf("Пользователь найден по токену %v", err)
	}
	if request == 2 {
		t.Errorf("Пользователь найден по токену  %v", sessionStorage.Sessions[0].AuthToken)
	}
}

func TestGetOnlineUserIds(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 2, AuthToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, AuthToken: "713e50a0651541d9b973aba3ec04e1f1"})
	userid := []uint{2, 3}
	request, err := sessionStorage.GetOnlineUserIds()
	if reflect.DeepEqual(request, userid) != true {
		t.Errorf("Не верный список пользователей %v %v", request, userid)
	}
	if err != nil {
		t.Errorf("Не получили список онлайн пользователей %v", err)
	}
}
func TestDeleteSessionByIndex(t *testing.T) {

	sessionStorage := new(SessionStorageMemory)
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 1, AuthToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 2, AuthToken: "713e50a0651541d9b973aba3ec04e1f1"})

	outSessionStorage := new(SessionStorageMemory)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 1, UserId: 1, AuthToken: "a396776f58b942fb9b10ebc798ab6303"})
	request := deleteSessionByIndex(sessionStorage.Sessions, 1)

	if reflect.DeepEqual(request, outSessionStorage.Sessions) != true {
		t.Errorf("\nНе верно удалилась сессия\n%v\n%v", request, outSessionStorage.Sessions)
	}

}
func TestGetLastSessionIdEmpty(t *testing.T) {
	sessionStorage := new(SessionStorageMemory)
	if getLastSessionId(sessionStorage) != 0 {
		t.Errorf("Данные о сессиях пресутсвуют")
	}
}
func TestGetLastSessionId(t *testing.T) {

	sessionStorage := new(SessionStorageMemory)
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, AuthToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, AuthToken: "713e50a0651541d9b973aba3ec04e1f1"})
	if getLastSessionId(sessionStorage) != sessionStorage.Sessions[0].ID {
		t.Errorf("Не верное выводиться index последней сессии")
	}
}

func TestDeleteSession(t *testing.T) {
	sessionStorage := new(SessionStorageMemory)
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, AuthToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, AuthToken: "713e50a0651541d9b973aba3ec04e1f1"})
	outSessionStorage := new(SessionStorageMemory)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, AuthToken: "713e50a0651541d9b973aba3ec04e1f1"})

	if sessionStorage.Delete(sessionStorage.Sessions[0].AuthToken) != nil {
		t.Errorf("Сессия не удалилась")
	}
}
func TestCreateSession(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, AuthToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, AuthToken: "713e50a0651541d9b973aba3ec04e1f1"})
	outSessionStorage := new(SessionStorageMemory)
	request, err := sessionStorage.Create(4)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 5, UserId: 4, AuthToken: "4eadd229ce654553a9b2a8fd13efd00"})
	if request == outSessionStorage.Sessions[0] {
		t.Errorf("токены сессий одинаковые %v %v", outSessionStorage.Sessions[0], request)
	}
	if err != nil {
		t.Errorf("Не создалась сессия %v", err)
	}

}
func TestISessionStorage(t *testing.T) {
	inter := NewSessionStorageMemory(new(TokenGeneratorUUID))

	session, err := inter.Create(1)
	if err != nil {
		t.Errorf("Не создалась сессия %v", err)
	}
	_, err = inter.GetOnlineUserIds()
	if err != nil {
		t.Errorf("Не получили список онлайн пользователей %v", err)
	}
	err = inter.Delete(session.AuthToken)
	if err != nil {
		t.Errorf("Не получили список онлайн пользователей %v", err)
	}

}
