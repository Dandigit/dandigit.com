// Start constant variables
var

// Grid vars
COLS=10,
ROWS=7,

// Canvas vars
EMPTY=0,
CHAR=1,
OBST=2,

// Control vars
LEFT=0,
UP=1,
RIGHT=2,
DOWN=3,
KEY_LEFT=37,
KEY_UP=38,
KEY_RIGHT=39,
KEY_DOWN=40,
// End constant variables

// Start gameobject variables
canvas, // HTML canvas element
ctx, // Canvas context
keystate, // Object that checks for k/b input
frames, // Used for animation
score; // Keeps track of the score for the very popular scoreboard
// End variables

// Structuring the grid - useful as this game uses a grid to render position of objects
grid = {
	width: null; // Defining grid width
	height: null; // Defining grid height
	_grid: null; // Defining grid array

	// Initiate the grid
	init: function(d, c, r) { // d = default value to fill
		this.width = c; // c = number of columns
		this.height = r; // r = number of rows
		this._grid = [];
	}
}