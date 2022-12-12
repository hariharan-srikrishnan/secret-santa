package permute

import (
	"fmt"
	"testing"
	"santa.com/secret/models"
)

func TestGetPlayerMap(t *testing.T) {

	players := []*models.Player{}

	players = append(players, models.CreatePlayer("Aravindh", "Aravindh_Placeholder_Email"))
	players = append(players, models.CreatePlayer("Gundh", "harivader008@gmail.com"))
	players = append(players, models.CreatePlayer("Prasad", "mailtoprasadkrp@gmail.com"))
	players = append(players, models.CreatePlayer("Bob", "hariharan.srikrishnan123@gmail.com"))
	players = append(players, models.CreatePlayer("Sriram", "vsriram1012@gmail.com"))
	players = append(players, models.CreatePlayer("Pranav", "pranavfornow@gmail.com"))
	players = append(players, models.CreatePlayer("Prathiba", "prathisesh@gmail.com"))
	players = append(players, models.CreatePlayer("Vivek", "satsvivek@gmail.com"))

	m := GetPlayerMap(players)
	fmt.Printf("%v", GetPlayerMap(players))

	for k, v := range m {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := GetPlayerMap(tt.args.players); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("GetPlayerMap() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func Test_getDerangement(t *testing.T) {
	a := getDerangements(8)
	if (len(a) != 14833) {
		t.Fatalf("error in derangement calculation %d", len(a))
	} 
	t.Fatalf("")
}
