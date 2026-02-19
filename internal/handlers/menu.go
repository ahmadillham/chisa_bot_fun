package handlers

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"

	"chisa_bot/pkg/utils"
)

// MenuHandler handles the menu command.
type MenuHandler struct{}

// NewMenuHandler creates a new MenuHandler.
func NewMenuHandler() *MenuHandler {
	return &MenuHandler{}
}

// HandleMenu sends a list of all available commands.
func (h *MenuHandler) HandleMenu(client *whatsmeow.Client, evt *events.Message) {
	menu := `📋 *Daftar Perintah*
Prefix: . ! /

━━━ 🎮 *Games* ━━━
• .tebakkata
  _Susun kata acak menjadi benar_
• .tebakibukota
  _Tebak ibu kota negara_
• .tebaknegara
  _Tebak negara dari clue_
• .tebakbenda
  _Tebak benda dari clue_
• .tebakbendera
  _Tebak nama negara dari bendera_
• .tebakangka
  _Tebak angka 1-100 (Higher/Lower)_
• .kuis
  _Kuis pengetahuan umum_
• .nyerah / .skip
  _Menyerah / lewati pertanyaan_
• .leaderboard / .lb
  _Cek klasemen mingguan_

━━━ 🖼️ *Media* ━━━
• .sticker (.s)
  _Ubah gambar/video/GIF jadi sticker_
• .toimg
  _Ubah sticker jadi gambar_
• .showimg (.rv)
  _Ambil media View Once (Reply pesan)_

━━━ 📥 *Downloader* ━━━
• .dl <link>
  _Download IG, TikTok, FB, YouTube_
• .mp3 <link>
  _Download Audio (YouTube/TikTok)_

━━━ 👥 *Grup* ━━━
• .tagall
  _Mention semua anggota (Admin only)_
• .kick <member>
  _Kick member (Admin only)_
  
━━━ 🎮 *Fun* ━━━
• .cekkhodam <nama>
  _Cek khodam kamu_
• .cekjodoh <nama1> <nama2>
  _Cek kecocokan jodoh_
• .kerangajaib <tanya>
  _Tanya kerang ajaib_
• .siapadia <tanya>
  _Random pick anggota grup_
• .rate <sesuatu>
  _Rating random 0-100_
• .roast <nama>
  _Roasting lucu_
• .seberapa <sifat> <nama>
  _Seberapa X kamu?_


━━━ 🛠️ *Lainnya* ━━━
• .short <link>
  _Pendekkan link (TinyURL)_
• .pick <opsi1> | <opsi2>
  _Pilih opsi random_
• .stats
  _Status server bot_
• .menu
  _Tampilkan pesan ini_`

	utils.ReplyText(client, evt, menu)
}
