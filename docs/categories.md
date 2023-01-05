# Categories

## endpoint

- path
```
POST "/categories/add"
```

- body
``` json
{
    'categories': ['x', 'y', 'z']
}
```

## fluxo
* Seleciona as categorias do `body` da mensagem
* Verifica os nomes que já existem
* Seleciona novos
  * Procura playlists com a nova categoria [1]
  * Extrai todas as músicas da playlist [2]
  * Adiciona a nova categoria com as musicas extraídas
* retorna as categorias

## Tarefas para realizar o fluxo
[1] função para coletar playlists
[2] função para coletar músicas da playlist
