// transformer_test.go
package transformer

import (
	"encoding/json"
	"testing" // Import package testing bawaan Go
)

// TestTransformMessage adalah fungsi unit test untuk fungsi TransformMessage.
// Nama fungsi test harus diawali dengan "Test" dan menerima *testing.T sebagai argumen.
func TestTransformMessage(t *testing.T) {
	// Definisikan kasus uji (test cases)
	tests := []struct {
		name          string // Nama kasus uji
		inputJSON     []byte // Input JSON untuk fungsi TransformMessage
		expectedJSON  []byte // Output JSON yang diharapkan
		expectedError bool   // Apakah kita mengharapkan error
	}{
		{
			name:          "Sukses Transformasi Normal",
			inputJSON:     []byte(`{"id": "tx123", "payload": "halo dunia", "status": "PENDING"}`),
			expectedJSON:  []byte(`{"id":"tx123","payload":"HALO DUNIA","status":"PROCESSED"}`),
			expectedError: false,
		},
		{
			name:          "Sukses Transformasi dengan Payload Kosong",
			inputJSON:     []byte(`{"id": "tx456", "payload": "", "status": "NEW"}`),
			expectedJSON:  []byte(`{"id":"tx456","payload":"","status":"PROCESSED"}`),
			expectedError: false,
		},
		{
			name:          "Input JSON Invalid",
			inputJSON:     []byte(`{"id": "tx789", "payload": "invalid`), // JSON tidak valid
			expectedJSON:  nil,
			expectedError: true,
		},
		{
			name:          "Sukses Transformasi dengan Angka di Payload",
			inputJSON:     []byte(`{"id": "tx101", "payload": "data123", "status": "PENDING"}`),
			expectedJSON:  []byte(`{"id":"tx101","payload":"DATA123","status":"PROCESSED"}`),
			expectedError: false,
		},
	}

	// Iterasi melalui setiap kasus uji
	for _, tt := range tests {
		// t.Run memungkinkan pengujian sub-test, mempermudah identifikasi kegagalan
		t.Run(tt.name, func(t *testing.T) {
			// Panggil fungsi yang akan diuji
			actualOutput, err := TransformMessage(tt.inputJSON)

			// Periksa apakah ada error yang tidak diharapkan atau error yang diharapkan tidak muncul
			if (err != nil) != tt.expectedError {
				t.Fatalf("TransformMessage() error = %v, expectedError %v", err, tt.expectedError)
			}

			// Jika tidak ada error yang diharapkan, bandingkan output
			if !tt.expectedError {
				// Untuk membandingkan JSON, lebih baik deserialisasi dan bandingkan struct
				// Ini menghindari masalah urutan kunci dalam JSON string
				var actualMsg, expectedMsg Message
				json.Unmarshal(actualOutput, &actualMsg)
				json.Unmarshal(tt.expectedJSON, &expectedMsg)

				if actualMsg.ID != expectedMsg.ID || actualMsg.Payload != expectedMsg.Payload || actualMsg.Status != expectedMsg.Status {
					t.Errorf("TransformMessage() got = %s, want %s", actualOutput, tt.expectedJSON)
				}
			}
		})
	}
}

// Catatan:
// 1. Fungsi test harus berada di file dengan akhiran _test.go (misalnya transformer_test.go)
// 2. Fungsi test harus berada di package yang sama dengan kode yang diuji (atau package _test terpisah)
// 3. Gunakan `t.Errorf` untuk menandai kegagalan test tanpa menghentikan eksekusi sub-test lainnya.
// 4. Gunakan `t.Fatalf` untuk menandai kegagalan fatal yang harus menghentikan eksekusi sub-test.
// 5. `json.Unmarshal` dan `json.Marshal` digunakan di sini untuk mensimulasikan serialisasi/deserialisasi pesan Kafka.
//    Dalam pengujian unit, kita fokus pada logika transformasi, bukan interaksi Kafka sebenarnya.
