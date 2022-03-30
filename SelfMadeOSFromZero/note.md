# ゼロからのOS自作入門



## 第1章 PCの仕組みとハローワールド

### 1.1 ハローワールド
- バイナリファイルをバイナリエディタを使用して作成し、ブートローダ経由で起動するプログラムを作成する。
- セキュアブートの設定が有効である場合は、メモリ経由のブートが出来ないため設定を変更する。
```shell
$sum BOOTX64.EFI
12430
$
$sudo umount /dev/sdb1
$sudo mkfs.fat /dev/sdb1
$sudo mkdir -p /mnt/usbmem
$sudo mount /dev/sdb1 /mnt/usbmem
$sudo mkdir -p /mnt/usbmem/EFI/BOOT
$sudo cp BOOTX64.EFI /mnt/usbmem/EFI/BOOT
$sudo umount /mnt/usbmem

```
### 1.2 USBメモリのデバイス名の探し方
- USBメモリのデバイス名の検索は、dmesgを使用する。
```shell
$dmesg
```
### 1.3 WSLのやり方

### 1.4 エミュレータでのやり方

### 1.5 結局何をやったのか
- 実行可能ファイルをバイナリエディタで作成し、USBメモリに指定された名前で保存する事でパソコンに内臓されているUEFI BIOSがそのファイルを読み取って実行した。

### 1.6 とにかく手を動かそう
- まずは手を動かして写経した後、改造を加える事で理解を深めるのが効率が良いと思われる。

### 1.7 UEFI BIOSによる起動
- BIOS(Basic Input Output System)は、コンピュータの電源をいれて最初に実行されるファームウェアと呼ばれるプログラムである。
- BIOSはCPUに付属しているROMに格納されている。
- UEFIと呼ばれる標準仕様にしたがって作られたBIOSをUEFI BIOSという。またUEFI BIOSが読み込んで実行できるプログラムの事をUEFIアプリケーションと呼ぶ。
- BIOSが実行される流れ
    1. CPUによるBIOSの読み込みと実行
    2. BIOSによるコンピュータ・周辺機器の初期化
    3. BIOSによる実行可能ファイルのストレージ探索
    4. BIOSによる実行可能ファイルのメモリへの読み出し
    5. BIOSの実行中断
    6. 読み出した実行可能ファイルの実行開始

### 1.8 OSを作る道具
- 本書ではOSをC/C++を使用して作成する。
- C/C++はソースコードをコンパイルし、オブジェクトファイル化した後、リンカを使用して完全な実行可能ファイル(=実行可能バイナリ)へと変換する。

### 1.9 C言語でハローワールド
- 
```c
EFI_STATUS EfiMain(EFI_HANDLE ImageHandle, EFI_SYSTEM_TABLE *SystemTable){
    SystemTable->ConOut->OutString(SystemTable->ConOut, L"Hello, world!\n");
    while(1);
    return 0;
}
```


## 第2章 EDK II入門とメモリマップ
### 2.1 EDK II入門
- EDK IIは、IntelがEFIとその周辺のプログラムを実装し、それが後にオープンソースとして公開されたものである。
- UEFI BIOS自体の開発やUEFI BIOS上で動くアプリケーションの開発にも使う事が出来る。

### 2.2 EDK IIでハローワールド
- フォルダ構成は以下の通りである。
    - MikanLoader.dec
        - パッケージ宣言ファイル
    - MikanLoader.dsc
        - パッケージ記述ファイル
    - Loader.inf
        - コンポーネント定義ファイル
        ```
        [Defines]
            ENTRY_POINT         = UefiMain
        ```
    - Main.c
        - ソースコード
        ```
        #include <Uefi.h>
        #include <Library/UefiLib.h> #UEDK IIのインポート

        EFI_STATUS EFIAPI UefiMain(EFI_HANDLE image_handle, EFI_SYSYTEM_TABLE *system_table){
            Print(L"Hello, Mikan World!\n");
            while(1);
            return EFI_SUCCESS;
        }
        ```

### 2.3 メインメモリ
- メインメモリは、1バイトごとにアドレスが割り当てられている。

### 2.4 メモリマップ
- メモリマップは、メインメモリのどの部分がどの用途で使用されているかが乗っているマップである。
- 以下のプロパティを含む構造体である。
    - PhysicalStart
        - 最も先頭にあるメモリ領域を指す。
    - Type
        - その領域が何に使われているかを示している。
        - メモリマップの例
            - EfiLoaderCode
                - UEFIアプリケーションの実行コード
            - EfiLoaderData
                - UEFIアプリケーションが使うデータ領域
            - EfiBootServicesCode
                - ブートサービスドライバの実行コード
            - EfiBootServiesCode
                - ブートサービスドライバが使うデータ領域
            - EfiConversionMemory
                - 空き領域
    - NumberOfPages
        - UEFIのメモリマップは、1ページ4KBとなっている。

### 2.5 メモリマップの取得
### 2.6 メモリマップのファイルへの保存
### 2.7 メモリマップの確認
### 2.8 ポインタ入門(1): アドレスとポインタ
### 2.9 ポインタとアロー演算子


## 第3章 画面表示の練習とブートローダ
