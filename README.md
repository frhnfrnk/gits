## Link doc postman : https://documenter.getpostman.com/view/23670076/2s9YsFFEZM
[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://god.gw.postman.com/run-collection/23670076-291eb1d9-5737-460b-a802-041f3fc16be1?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D23670076-291eb1d9-5737-460b-a802-041f3fc16be1%26entityType%3Dcollection%26workspaceId%3D60c92e9e-e77c-4a32-8bd1-d67d468e2acb)

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
