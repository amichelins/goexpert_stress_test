package stress

import "time"

// Resposta Struct com os dados da resposta de cada chamada realizada.
// -------------------------------------------------------------------
//
//      Ini  time.Time Tempo de inicio da chamada
//
//      Fim  time.Time Tempo final da chamada
//
//      Code int Código http da reposta ( 2xx, 4xx, 5xx)
//
//      Erro error Erro na chamada
//
type Resposta struct {
    Ini  time.Time
    Fim  time.Time
    Code int
    Erro error
}

type Relatorio struct {
    RequisicoesFeitas  int64
    RequisicoesOk      int64
    Concorrencia       int64
    TempoGasto         string
    InfTotais          int64
    RedirecoesTotais   int64
    ProblemaaCliTotais int64
    PorblemasSrvTotais int64
    Erros              int64
}
