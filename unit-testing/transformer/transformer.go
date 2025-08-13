// transformer.go
package transformer

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Message adalah struktur untuk merepresentasikan pesan Kafka.
// Ini adalah bagian dari 'Model' atau 'Data' concern.
type Message struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
	Status  string `json:"status"`
}

// TransformMessage adalah fungsi yang mengimplementasikan logika transformasi pesan.
// Ini adalah bagian dari 'Logic' atau 'Service' concern.
// Fungsi ini menerima input JSON (yang mensimulasikan payload pesan Kafka),
// melakukan transformasi (misalnya, mengubah teks menjadi huruf kapital),
// dan mengembalikan pesan yang sudah ditransformasi dalam format JSON.
func TransformMessage(inputJSON []byte) ([]byte, error) {
	var msg Message
	// Deserialisasi input JSON ke struct Message
	err := json.Unmarshal(inputJSON, &msg)
	if err != nil {
		return nil, fmt.Errorf("gagal deserialisasi pesan: %w", err)
	}

	// Lakukan transformasi: ubah payload menjadi huruf kapital
	msg.Payload = strings.ToUpper(msg.Payload)
	msg.Status = "PROCESSED" // Ubah status setelah diproses

	// Serialisasi kembali struct Message ke JSON
	outputJSON, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("gagal serialisasi pesan: %w", err)
	}

	return outputJSON, nil
}

// Catatan: Dalam aplikasi Kafka sungguhan, fungsi ini akan dipanggil oleh consumer
// setelah menerima pesan dari topic, dan hasilnya mungkin akan dipublikasikan
// ke topic lain oleh producer. Untuk unit test, kita hanya fokus pada logika TransformMessage.
