package firsttask

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const separator = "------------------------"

func Run(in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)

	fmt.Fprintln(out, "Aplikasi Kasir Toko Kelontong")
	fmt.Fprintln(out, "Ketik 'batal' pada input harga barang untuk keluar.")
	fmt.Fprintln(out)

	for {
		itemPrice, ok, err := readAmount(reader, out, "Harga Barang: ")
		if err != nil {
			return err
		}
		if !ok {
			fmt.Fprintln(out, "Sampai jumpa!")
			return nil
		}

		paidAmount, ok, err := readAmount(reader, out, "Uang Pembeli: ")
		if err != nil {
			return err
		}
		if !ok {
			fmt.Fprintln(out, "Sampai jumpa!")
			return nil
		}

		printResult(out, itemPrice, paidAmount)
		fmt.Fprintln(out, separator)
	}
}

func printResult(out io.Writer, itemPrice, paidAmount int64) {
	tx, err := newTransaction(itemPrice, paidAmount)
	if err != nil {
		fmt.Fprintf(out, "[SISTEM] %s\n", err)
		return
	}

	r := tx.evaluate()
	if !r.approved {
		fmt.Fprintf(out, "[SISTEM] Transaksi Ditolak! Uang kurang %d.\n", r.shortfall)
		return
	}

	fmt.Fprintf(out, "[SISTEM] Transaksi Berhasil. Kembalian Anda: %d.\n", r.change)
}

func readAmount(reader *bufio.Reader, out io.Writer, prompt string) (int64, bool, error) {
	for {
		fmt.Fprint(out, prompt)

		line, readErr := reader.ReadString('\n')
		if readErr != nil && readErr != io.EOF {
			return 0, false, readErr
		}

		input := strings.TrimSpace(line)
		if input == "" && readErr == io.EOF {
			return 0, false, nil
		}
		if strings.EqualFold(input, "batal") {
			return 0, false, nil
		}

		amount, convErr := strconv.ParseInt(input, 10, 64)
		if convErr != nil {
			fmt.Fprintln(out, "[SISTEM] Input tidak valid, masukkan angka bulat.")
			if readErr == io.EOF {
				return 0, false, nil
			}
			continue
		}

		return amount, true, nil
	}
}
