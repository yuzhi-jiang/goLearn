package compose

// 坐标点point点
type point struct {
	x int
	y int
}

// 方向
type direction int

const (
	Up direction = iota
	Down
	Left
	Right
)

type snake struct {
	body []point
	dir  direction
	len  int
}

func (s *snake) setSnakeBody(snakeBody []point) {
	s.body = snakeBody
}

// set Direction
func (s *snake) setDirection(dir direction) {
	s.dir = dir
}

// get body
func (s *snake) getSnakeBody() []point {
	return s.body
}

// get direction
func (s *snake) getDirection() direction {
	return s.dir
}

// get len
func (s *snake) getLen() int {
	return s.len
}

// /////////屏幕
type screen struct {
	width  int
	height int
	score  int
	snakes *snake
	foods  *point
}

const (
	WIDTH  = 60
	HEIGHT = 40
)

// new screen
func newScreen() *screen {
	return &screen{snakes: new(snake), foods: new(point), width: WIDTH, height: HEIGHT, score: 0}
}

// get score from screen
func (s *screen) getScore() *int {
	return &s.score
}

// set score
func (s *screen) setScore(score int) {
	s.score = score
}

// 初始化screen
func (s *screen) initScreenSize() {
	s.height = HEIGHT
	s.width = WIDTH
}

// get snake from screen
func (s *screen) getSnake() *snake {
	return s.snakes
}

// get food from screen
func (s *screen) getFood() *point {
	return s.foods
}

// get width from screen
func (s *screen) getWidth() int {
	return s.width
}

// get height from screen
func (s *screen) getHeight() int {
	return s.height
}

// ///////////控制器
type status int

const (
	//暂停
	PAUSE status = iota

	RUN

	GAMEOVER
)

type controller struct {
	//snake 状态 通道
	snakeStatusChannel chan bool
	//游戏状态
	gameStatusChannel chan status
	//方向
	dir direction
}

// 实例化controller
func newController() *controller {
	return &controller{
		snakeStatusChannel: make(chan bool, 1),
		gameStatusChannel:  make(chan status, 1),
		dir:                0,
	}
}

// set gameStatusChannel for controller
func (c *controller) setGameStatusChannel(gameStatusChannel chan status) {
	c.gameStatusChannel = gameStatusChannel
}

// get gameStatusChannel for controller
func (c *controller) getGameStatusChannel() chan status {
	return c.gameStatusChannel
}

// get snakeStatusChannel for controller
func (c *controller) getSnakeStatusChannel() chan bool {
	return c.snakeStatusChannel
}

// set snakeStatusChannel for controller
func (c *controller) setSnakeStatusChannel(snakeStatusChannel chan bool) {
	c.snakeStatusChannel = snakeStatusChannel
}

// get dir
func (c *controller) getDir() direction {
	return c.dir
}

// //////游戏总览
type game struct {
	//控制
	controller *controller
	//屏幕
	screen *screen
}

// get controller from game
func (g *game) getController() *controller {
	return g.controller
}

// get screen from game
func (g *game) getScreen() *screen {
	return g.screen
}

// 实例化 game
func newGame() *game {
	return &game{
		controller: newController(),
		screen:     newScreen(),
	}
}
