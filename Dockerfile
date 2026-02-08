# Gunakan Go versi 1.20 yang stabil
FROM golang:1.20

# Buat folder kerja di dalam server
WORKDIR /app

# Copy semua file kode Anda ke dalam server
COPY . .

# Download library yang dibutuhkan (otomatis)
RUN go mod download

# Rakit (Build) kodenya jadi aplikasi bernama 'server'
RUN go build -o server main.go

# Jalankan aplikasinya
CMD ["./server"]
