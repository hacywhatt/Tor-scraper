# Read Me
Toolu kullanamk için öncelikle go modülünü bilgisayarınıza indirmeniz gerekmektedir. https://go.dev/dl/ bu link üzerinden indirebilirsiniz. <br>
Daha sonra modülü başlatmak için <br>  
cd web_scrapper<br>
go mod init web_scrapper <br>
komutları girilmelidir.<br>
Eğer bilgisayarınızda gerekli kütüphane yoksa<br>
go get golang.org/x/net/proxy<br>
komutuyla indirebilirsiniz. <br>

Toolun düzgün çalışması için, sisteminizde Tor kurulu ve bağlantısı yapılmış şekilde olmalıdır. Ayrıca hedef.yaml dosyası olması gerekmektedir.<br>


Toolu çalıştırmak için CMD  ekranında <br>
go run main.go  <br>
komutunu girerek çalıştırabilirsiniz.

NOT: kodun default ayarlarında bir windows makinede çalıştırıldığı varsayılarak MSEDGE yolu   C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe  olarak verilmiştir. Eğer MSEDGE yolunun farklı olduğunu düşünüyorsanız 125. satırdaki kodu değiştirmeniz gerekmektedir.
