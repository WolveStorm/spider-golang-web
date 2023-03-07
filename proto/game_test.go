package proto

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func TestGameList(t *testing.T) {
	dial, err := grpc.Dial(":8770", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	client := NewGameClient(dial)
	list, err := client.GameList(context.Background(), &GameListFilterRequest{
		Page:     1,
		PageSize: 5,
		Keyword:  "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(list)
}

func TestGameDetail(t *testing.T) {
	dial, err := grpc.Dial(":8770", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	client := NewGameClient(dial)
	list, err := client.GameDetail(context.Background(), &GameDetailRequest{
		GameName: "Roblox",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(list)
}
