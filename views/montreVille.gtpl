<html>

<head>
    <link rel="stylesheet" href="/assets/ville.css" type="text/css">
    </link>
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Josefin+Slab">

    <title>{{ .Name }}</title>
</head>

<body>
    {{ if .Name}}
    <div class="infos-ville">
        <h1>{{ .Owner }} est le maire de <span id="nom-ville">{{ .Name }}</span></h1>
        <h2>Progrès technologique : <span id="points-techno"> {{ .Progression}} </span> </h2>
    </div>
    {{else}}
    <div class="infos-ville">
        <h1>Aucun résultat pour cette combinaison <span id="points-techno">Maire</span> / <span
                id="nom-ville">Ville</span></h1>
        <h2><a href="/annuaire">Retour à la recherche</a></h2>
    </div>
    {{end}}
    <br>
    <br>
    <br>
    <div class="river">
        <div class='wave'></div>
        <div class='wave'></div>
        <div class='wave'></div>
        <div class='wave'></div>
        <div class='wave'></div>
        <div class='wave'></div>
        <div class='wave'></div>
    </div>
    <br>
    <br>
    <br>
    <div class="ui-wrapper">
        <div class="canvas-screen">
            <canvas id="c" style="  position:absolute; top:27%; left:30%; z-index:0" width="704" height="704"></canvas>
            <canvas id="d" style=" position:absolute; top:27%; left:30%; z-index:1; pointer-events:none;" width="704"
                height="704"></canvas>
        </div>
        <div style="padding-left: 724px;">
            <ul style="list-style-type: none;">
                <li class="build-select"> <img id="rochers" data-atlas-index="13" src="/assets/tiles/rochers_duo.png">
                    <h3>Rochers</h3>
                </li>
                <li class="build-select"> <img id="arbres" data-atlas-index="1" src="/assets/tiles/arbres_duo.png">
                    <h3>Arbres </h3>
                </li>
            </ul>
        </div>

    </div>
</body>
<script type="text/javascript" src="/assets/ville.js"></script>

</html>