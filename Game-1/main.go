package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running   = true
	bckgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D
	music        rl.Music

	cam           rl.Camera2D
	playerSrc     rl.Rectangle
	playerDest    rl.Rectangle
	playerSpeed   float32 = 4
	isMsuicPaused bool
)

func iinit() {
	rl.InitWindow(screenWidth, screenHeight, "졸리크의 게임")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	music = rl.LoadMusicStream("res/musiq.mp3")
	grassSprite = rl.LoadTexture("res/ts/Grass.png")
	playerSprite = rl.LoadTexture("res/chars/char.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(float32(playerDest.X-playerDest.Width/2),
			float32(playerDest.Y-playerDest.Height/2)), 0.0, 1.0)

	rl.InitAudioDevice()
	rl.PlayMusicStream(music)
}

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		isMsuicPaused = !isMsuicPaused
	}
}
func update() {
	running = !rl.WindowShouldClose()

	rl.UpdateMusicStream(music)
	if isMsuicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-playerDest.Width/2), float32(playerDest.Y-playerDest.Height/2))

	fmt.Printf("X = %v\n", playerDest.X)
	fmt.Printf("Y = %v\n", playerDest.Y)
}
func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bckgColor)
	rl.BeginMode2D(cam)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.Black)
	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}
func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {
	iinit()
	for running {
		input()
		update()
		render()
	}
	quit()
}
