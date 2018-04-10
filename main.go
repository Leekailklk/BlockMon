// curl -X POST http://101.37.164.153:9888/get-block -d '{"block_height": 26}'
package main

import(
    "log"
    // "fmt"
    "strconv"
    "time"
    "encoding/json"
    "io/ioutil"
    
    "github.com/parnurzeal/gorequest"
    // "github.com/bytom/consensus/difficulty"
)


type t_data struct {
    BlockCount  uint64  `json:"block_count, omitempty"`
    Hash        string  `json:"hash, omitempty"`
    PreBlckHsh  string  `json:"previous_block_hash, omitempty"`
    Size        uint64  `json:"size, omitempty"`
    Version     uint8   `json:"version, omitempty"`
    Height      uint64  `json:"height, omitempty"`
    Timestamp   int64  `json:"timestamp, omitempty"`
    Nonce       uint64  `json:"nonce, omitempty"`
    Bits        uint64  `json:"bits, omitempty"`
    Diff        string  `json:"difficulty, omitempty"`
}

type t_resp struct {
    Status      string  `json:"status"`
    Data        t_data  `json:"data"`
}


const (
    walletAddr = "http://101.37.164.153:9888/"
)


func main() {
    request := gorequest.New()
    var resp t_resp

    _, body, _ := request.Post(walletAddr + "get-block-count").
        End()

    json.Unmarshal([]byte(body), &resp)
    if resp.Status != "success" {
        log.Fatalln("Request fail!")    
    }
    log.Printf("Block Count: %d\n\n", resp.Data.BlockCount)

    var dataStr string
    for i := uint64(1); i <= resp.Data.BlockCount; i++ {
        _, body, _ = request.Post(walletAddr + "get-block").
            Send(`{
                    "block_height": `+ strconv.FormatUint(i, 10) + `
                    }`).
            End()
        json.Unmarshal([]byte(body), &resp)
        // log.Printf("Block %d of %d:\n\tTimestamp: %d, %v\n\tNonce: %d\n\tBits: %d, e.g., %d\n\tDiff: %s",
        // log.Printf("Block %d of %d:\n\tTimestamp: %d, %v\n\tNonce: %d\n\tBits: %d\n\tDiff: %s",
        log.Printf("Block %d of %d:\n\tTimestamp: %d, %v\n\tDiffi: %s",
            resp.Data.Height, resp.Data.BlockCount,
            resp.Data.Timestamp, time.Unix(resp.Data.Timestamp,0),
            // resp.Data.Nonce,
            // resp.Data.Bits, difficulty.CompactToBig(resp.Data.Bits),
            // resp.Data.Bits,
            resp.Data.Diff)

            dataStr += strconv.FormatUint(resp.Data.Height, 10)
            dataStr += ","
            dataStr += strconv.FormatInt(resp.Data.Timestamp, 10)
            dataStr += ","
            dataStr += resp.Data.Diff
            dataStr += ",("
            dataStr += time.Unix(resp.Data.Timestamp,0).String()
            dataStr += "),\n"
    }
    err := ioutil.WriteFile("./all-blocks.csv", []byte(dataStr), 0644)
    check(err)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
