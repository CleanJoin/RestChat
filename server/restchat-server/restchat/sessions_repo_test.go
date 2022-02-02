package restchat

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetOnlineUserIds(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 1, Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 2, Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})
	userid := []int{1, 2}
	request, err := sessionStorage.GetOnlineUserIds()
	if reflect.DeepEqual(request, userid) != false {
		t.Errorf("Не верный список пользователей")
	}

	fmt.Printf("Все хорошо! массив готов %v\n", err)
}
func TestDeleteSessionByIndex(t *testing.T) {

	sessionStorage := new(SessionStorageMemory)
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 1, UserId: 1, Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 2, Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})

	outSessionStorage := new(SessionStorageMemory)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 1, UserId: 1, Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	request := deleteSessionByIndex(sessionStorage.Sessions, 1)

	if reflect.DeepEqual(request, outSessionStorage.Sessions) != true {
		t.Errorf("\nНе верно удалилась сессия\n%v\n%v", request, outSessionStorage.Sessions)
	}
	fmt.Printf("По указанному index удалили сессию %v\n", outSessionStorage.Sessions)
}
func TestGetLastSessionIdEmpty(t *testing.T) {
	sessionStorage := new(SessionStorageMemory)
	if getLastSessionId(sessionStorage) != 0 {
		t.Errorf("Данные о сессиях пресутсвуют")
	}
	fmt.Println("Нет данных о сессиях")
}
func TestGetLastSessionId(t *testing.T) {

	sessionStorage := new(SessionStorageMemory)
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})
	if getLastSessionId(sessionStorage) != sessionStorage.Sessions[0].ID {
		t.Errorf("Не верное выводиться index последней сессии")
	}
	fmt.Printf("Идентификатор последней сессии: %v\n", getLastSessionId(sessionStorage))
}

func TestDeleteSession(t *testing.T) {
	sessionStorage := new(SessionStorageMemory)
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})
	api_token := "a396776f58b942fb9b10ebc798ab6303"
	outSessionStorage := new(SessionStorageMemory)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})
	request, err := sessionStorage.Delete(api_token)

	if reflect.DeepEqual(request.Sessions, outSessionStorage.Sessions) != true {
		t.Errorf("\nНе верно удалилась сессия\n%v\n%v", request, outSessionStorage.Sessions)
	}
	fmt.Println(err)
}
func TestCreateSession(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 4, UserId: 1, Auth_token: "a396776f58b942fb9b10ebc798ab6303"})
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: 2, UserId: 3, Auth_token: "713e50a0651541d9b973aba3ec04e1f1"})
	outSessionStorage := new(SessionStorageMemory)
	request, _ := sessionStorage.Create(4)
	outSessionStorage.Sessions = append(outSessionStorage.Sessions, SessionModel{ID: 5, UserId: 4, Auth_token: "4eadd229ce654553a9b2a8fd13efd00"})
	if request == outSessionStorage.Sessions[0] {
		t.Errorf("токены сессий одинаковые %v %v", outSessionStorage.Sessions[0], request)
	}
	fmt.Printf("Сессия создана: %v", request)

}
