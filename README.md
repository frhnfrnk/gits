# Detail kompleksitas

Kompleksitas kodingan ini dapat dijelaskan dalam beberapa aspek, yaitu kompleksitas waktu (time complexity) dan kompleksitas ruang (space complexity).

1. **Kompleksitas Waktu:**
   - Iterasi melalui setiap karakter dalam input memiliki kompleksitas waktu O(n), di mana n adalah panjang input.
   - Operasi seperti `append` dan pemotongan slice memiliki kompleksitas O(1) pada rata-rata kasus, tetapi dapat menjadi O(n) pada kasus terburuk di mana alokasi ulang memerlukan kopi seluruh slice.
   - Karena setiap karakter diproses sekali dan hanya ada satu kali iterasi melalui input, kompleksitas waktu totalnya adalah O(n).

Sehingga, kompleksitas waktu dari fungsi `isBalanced` adalah O(n), di mana n adalah panjang input.

2. **Kompleksitas Ruang:**
   - Penggunaan stack menyebabkan penggunaan ruang tambahan proporsional terhadap panjang tumpukan (stack).
   - Pada kasus terburuk, tumpukan dapat mencapai ukuran n (panjang input) jika setiap karakter dalam input adalah tanda buka dan tidak ada tanda tutup yang sesuai. Jadi, kompleksitas ruang dalam kasus terburuk adalah O(n).
   - Selain itu, penggunaan peta (map) `bracketMap` dengan konstan jumlah elemen tidak memengaruhi kompleksitas ruang secara signifikan.

Sehingga, kompleksitas ruang dari fungsi `isBalanced` adalah O(n).

Dengan demikian, kodingan ini memiliki kompleksitas waktu O(n) dan kompleksitas ruang O(n) pada kasus terburuk.
