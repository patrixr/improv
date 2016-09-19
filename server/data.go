package main

import (
    "encoding/json"
    "math"
    "math/rand"
    "strconv"
    "log"
)

var CHUNK_LEN = 100

/**
 * Dataset
 * The representation of each chunk of timeline
 *
 */
type Dataset struct {
    Alpha   []int   `json:"alpha"`
    Beta    []int   `json:"beta"`
    Gamma   []int   `json:"gamma"`
    Delta   []int   `json:"delta"`
    Epsilon []int   `json:"epsilon"`
}

func (this *Dataset) Append(other *Dataset) {
    this.Alpha      = append(this.Alpha, other.Alpha...)
    this.Beta       = append(this.Beta, other.Beta...)
    this.Gamma      = append(this.Gamma, other.Gamma...)
    this.Delta      = append(this.Delta, other.Delta...)
    this.Epsilon    = append(this.Epsilon, other.Epsilon...)
}

func (this *Dataset) Slice(from int, to int) *Dataset {
    ds := Dataset{}
    ds.Alpha      = this.Alpha[from:to]
    ds.Beta       = this.Beta[from:to]
    ds.Gamma      = this.Gamma[from:to]
    ds.Delta      = this.Delta[from:to]
    ds.Epsilon    = this.Epsilon[from:to]
    return &ds
}

/**
 * ImprovStorage
 * Main Storage Class. Returns any position within the timeline
 *
 */
type ImprovStorage struct {
    // A Database extension
    Database
}

func (this *ImprovStorage) _generateLine() []int {
    var line []int  
    for i := 0; i < CHUNK_LEN; i++ {
        if rand.Intn(3) == 2 {
            line = append(line, rand.Intn(9))
        } else {
            line = append(line, 0)
        }
    }
    return line
}

func (this *ImprovStorage) _generateChunk() Dataset {
    return Dataset{
        Alpha:      this._generateLine(),
        Beta:       this._generateLine(),
        Gamma:      this._generateLine(),
        Delta:      this._generateLine(),
        Epsilon:    this._generateLine(),
    }
}

func (this *ImprovStorage) _getChunk(idx int) *Dataset {

    var chunk Dataset

    id      := "chunk_" + strconv.Itoa(idx)
    str     := this.Get(id)

    if str == "" {
        //
        // This chunk does not exist yet
        chunk = this._generateChunk()
        str, err := json.Marshal(&chunk)
        if err != nil { log.Panic("JSON.Marshal Failed") }
        this.Set(id, string(str))
    } else {
        //
        // We found the data
        err := json.Unmarshal([]byte(str), &chunk)
        if err != nil {
            log.Panic("JSON.Unmarshal Failed")
        }
    }
    return &chunk
}

func (this *ImprovStorage) Read(from int, count int) *Dataset {
    chunkIdx    := int(math.Floor( float64(from / CHUNK_LEN) ))
    startIdx    := from - chunkIdx * 100

    chunk       := this._getChunk(chunkIdx)
    chunkCount  := (count + startIdx) / CHUNK_LEN

    for i := 1; i <= chunkCount; i++ {
        next := this._getChunk(chunkIdx + i)
        chunk.Append(next)
    }

    return chunk.Slice(startIdx, startIdx + count)
}