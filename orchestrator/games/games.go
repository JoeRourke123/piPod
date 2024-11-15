package games

type Game struct {
	Title string
	Path  string
	Icon  string
}

var (
	ActiveGames = []Game{
		{
			Title: "Brickbreaker",
			Path:  "/game/brickbreaker",
			Icon:  "Wall",
		},
		{
			Title: "Snake",
			Path:  "/game/snake",
			Icon:  "Steps",
		},
		{
			Title: "2048",
			Path:  "/game/2048",
			Icon:  "SquaresFour",
		},
		{
			Title: "Wordle",
			Path:  "/game/wordle",
			Icon:  "DotsThreeCircle",
		},
		{
			Title: "Space Invaders",
			Path:  "/game/spaceinvaders",
			Icon:  "FlyingSaucer",
		},
	}
)
