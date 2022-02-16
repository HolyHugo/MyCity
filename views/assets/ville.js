var currentIndex = 1;
window.onload = function () {
    document.getElementById("rochers").addEventListener("click", changeNode);
    document.getElementById("arbres").addEventListener("click", changeNode);
    var map = {
        cols: 11,
        rows: 11,
        tsize: 64,
        tiles: [
            03, 03, 01, 03, 05, 06, 03, 03, 03, 03, 13,
            13, 03, 03, 03, 07, 08, 03, 03, 01, 03, 02,
            04, 03, 03, 09, 10, 03, 03, 03, 03, 03, 02,
            04, 03, 03, 11, 12, 03, 01, 03, 03, 03, 13,
            04, 04, 03, 03, 03, 03, 03, 03, 03, 02, 02,
            04, 13, 03, 03, 03, 03, 03, 03, 13, 02, 02,
            04, 03, 03, 03, 03, 03, 03, 03, 02, 02, 02,
            03, 03, 01, 03, 03, 03, 03, 03, 03, 02, 02,
            03, 03, 03, 03, 03, 03, 01, 03, 13, 02, 02,
            13, 03, 03, 03, 03, 03, 03, 03, 02, 02, 02,
            03, 03, 03, 03, 03, 03, 13, 02, 02, 02, 02,
        ],
        getTile: function (col, row) {
            return this.tiles[row * map.cols + col]
        }
    };

    var buildableSquare = [
        5, 6, 7, 8,
        16, 17, 18,
        27, 28, 29, 30,
        38,40, 41,
        49, 50, 51, 52,
        60, 61, 62,
        71, 72, 73,
        82,83,84
    ]
    var canvas = document.getElementById('c');
    var context = document.getElementById('c').getContext('2d');
    var canvasBuilding = document.getElementById('d');
    var contextBuilding = document.getElementById('d').getContext('2d');

    var BB = canvas.getBoundingClientRect();
    var offsetX = BB.left;
    var offsetY = BB.top;
    canvas.onmousedown = handleMouseDown;
    var tileAtlas = new Image();

    tileAtlas.onload = function () {
        renderInitialMap(map, this, context);
    };
    tileAtlas.src = "/assets/tiles/tileAtlas.png";


    function handleMouseDown(e) {
        if (e.button === 0) {
            e.preventDefault();
            e.stopPropagation();

            // calculate the mouse position
            var mouseX = e.clientX - offsetX;
            var mouseY = e.clientY - offsetY;

            // translate mouse pos into grid pos
            var coordX = Math.floor(mouseX / 64)
            var coordY = Math.floor(mouseY / 64)
            // retrieve id of the clicked square in array
            var clickedSquare = coordX - 11 * coordY;

            //if clicked square is in the buildable ones @TODO check if already building on pos
            if (buildableSquare.includes(clickedSquare)) {
                if (confirm('Construire sur cette case ? ' + coordX + ':' + coordY)) {
                    contextBuilding.drawImage(
                        tileAtlas, // image
                        (currentIndex - 1) * map.tsize, // source x
                        0, // source y
                        map.tsize, // source width
                        map.tsize, // source height
                        coordX * map.tsize,  // target x
                        coordY * map.tsize, // target y
                        map.tsize, // target width
                        map.tsize // target height
                    );
                }
            } else {
                alert('Pas constructible !')
            }
        }
    }
}

function renderInitialMap(map, tileAtlas, context) {
    for (var c = 0; c < map.cols; c++) {
        for (var r = 0; r < map.rows; r++) {
            var tile = map.getTile(c, r);
            if (tile !== 0) { // 0 => empty tile
                context.drawImage(
                    tileAtlas, // image
                    (tile - 1) * map.tsize, // source x
                    0, // source y
                    map.tsize, // source width
                    map.tsize, // source height
                    c * map.tsize,  // target x
                    r * map.tsize, // target y
                    map.tsize, // target width
                    map.tsize // target height
                );
            }
        }
    }
}

function changeNode() {
    currentIndex = this.dataset.atlasIndex;
}