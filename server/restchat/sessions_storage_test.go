package restchat

import (
	"reflect"
	"testing"
)

//test написать
func TestGetUserId(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))

	var inter ISessionStorage = sessionStorage

	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 0, UserId: 2, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})

	request, err := inter.GetUserId(sessionStorage.Sessions[0].ApiToken)
	if request != 2 {
		t.Errorf("Пользователь не найден по токену  %v", sessionStorage.Sessions[0].ApiToken)
	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetUserIdBadApiToken(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 2, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})

	request, err := sessionStorage.GetUserId("a396776f58b942fb9b10ebc798ab630")

	if err == nil {
		t.Errorf("Пользователь найден по токену %v", err)
	}
	if request == 2 {
		t.Errorf("Пользователь найден по токену  %v", sessionStorage.Sessions[0].ApiToken)
	}
}

func TestGetOnlineUserIds(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 2, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, ApiToken: "713e50a0651541d9b973aba3ec04e1f1"})
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
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 2, ApiToken: "713e50a0651541d9b973aba3ec04e1f1"})

	outSessionStorage := new(SessionStorageMemory)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 1, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
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
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, ApiToken: "713e50a0651541d9b973aba3ec04e1f1"})
	if getLastSessionId(sessionStorage) != sessionStorage.Sessions[0].ID {
		t.Errorf("Не верное выводиться index последней сессии")
	}
}

func TestDeleteSession(t *testing.T) {
	sessionStorage := new(SessionStorageMemory)
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, ApiToken: "713e50a0651541d9b973aba3ec04e1f1"})
	outSessionStorage := new(SessionStorageMemory)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, ApiToken: "713e50a0651541d9b973aba3ec04e1f1"})

	if sessionStorage.Delete(sessionStorage.Sessions[0].ApiToken) != nil {
		t.Errorf("Сессия не удалилась")
	}
}
func TestCreateSession(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))

	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, ApiToken: "713e50a0651541d9b973aba3ec04e1f1"})
	outSessionStorage := new(SessionStorageMemory)
	request, err := sessionStorage.Create(4)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 5, UserId: 4, ApiToken: "4eadd229ce654553a9b2a8fd13efd00"})
	if request == outSessionStorage.Sessions[0] {
		t.Errorf("токены сессий одинаковые %v %v", outSessionStorage.Sessions[0], request)
	}
	if err != nil {
		t.Errorf("Не создалась сессия %v", err)
	}

}

func TestCreateNewSession(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 11, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 12, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 21, UserId: 2, ApiToken: "713e50a0651541d9b973aba3ec04e1f1"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 13, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 14, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 15, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 22, UserId: 2, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 16, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 31, UserId: 3, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 32, UserId: 3, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 33, UserId: 3, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 34, UserId: 3, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 17, UserId: 1, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 35, UserId: 3, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 36, UserId: 3, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 41, UserId: 4, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 42, UserId: 4, ApiToken: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Create(1)
	sessionStorage.Create(2)
	sessionStorage.Create(3)
	sessionStorage.Create(4)

	if len(sessionStorage.Sessions) != 4 || len(sessionStorage.Sessions) < 4 {
		t.Errorf("Не корректно удались сессии")
	}
}
func TestISessionStorage(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))

	var inter ISessionStorage = sessionStorage

	session, err := inter.Create(1)
	if err != nil {
		t.Errorf("Не создалась сессия %v", err)
	}
	_, err = inter.GetOnlineUserIds()
	if err != nil {
		t.Errorf("Не получили список онлайн пользователей %v", err)
	}
	err = inter.Delete(session.ApiToken)
	if err != nil {
		t.Errorf("Не получили список онлайн пользователей %v", err)
	}

}
