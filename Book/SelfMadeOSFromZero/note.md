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
- UEFIの機能を使用して、メモリマップを取得する。
- UEFIには二つの機能が存在する。
    - ブートサービス
        - OSを起動するために必要な機能を提供するブートサービス
    - ランタイムサービス
        - OS起動前・OS起動後どちらでも使える機能を提供するランタイムサービス
```c
EFI_STATUS GetMemoryMap(struct MemoryMap* map){
    //マップのバッファがNULLだった場合エラー
    if (map->buffer == NULL){
        return EFI_BUFFER_TOO_SMALL;
    }

    map->map_size = map->buffer_size;
    //EFI_STATUS GetMemoryMap(
    //  IN OUT UINTN *MemoryMapSize,                書き込みメモリマップの大きさ
    //  IN OUT EFI_MEMORY_DESCRIPTOR *MemoryMap,    書き込みのメモリ領域の先頭ポインタ
    //  OUT UINTN *MapKey,                          メモリマップを識別する端の値を書き込む変数
    //  OUT UINTN *DescriptorSize                   ディスクリプタのサイズ
    //)
    // IN OUTはEDK II独自のマクロで関数の引数または戻り値どちらに使用されるかを示す。
    return gBS->GetMemoryMap(
        &map->map_size,
        (EFI_MEMORY_DESCRIPTOR*)map->buffer,
        &map->map_key,
        &map->descriptor_size,
        &map->descriptor_version);
    )
}
```

### 2.6 メモリマップのファイルへの保存
### 2.7 メモリマップの確認
### 2.8 ポインタ入門(1): アドレスとポインタ
### 2.9 ポインタとアロー演算子


## 第3章 画面表示の練習とブートローダ


### 3.1 QEMUモニタ
- ブートローダ作りを始める前にQEMUモニタを使用してデバッグする方法を記述する。
- レジスタの確認
    - ```(qemu)info registers```
- メモリメモリの中で指定したアドレス付近の値を表示する。
    - addrを先頭するメモリ領域の値を表示する。
    - /fmtのフォーマット
        - 個数 + n進数 + サイズ
            - 進数
                - x 16進数
                - d 10進数
                - i 機械語命令を逆アセンブリにする。
            - サイズ
                - b 1バイト
                - h 2バイト
                - w 4バイト
                - g 8バイト
    - ```x /4xb addr```

### 3.2 レジスタ
- CPUレジスタは高速にアクセスできる記憶領域である。
- 種類としては、汎用レジスタと特殊レジスタの二種類が存在する。
    - 汎用レジスタ
        - 一般の演算に使用する事が出来るレジスタ
        - 主な目的は値を記録する事にある。
        - RAX, RBX, RCX, RDX, RBP, RSI, RDI, RSP, R8～R15が該当する。
    - 特殊レジスタ
        - 実行に不可欠な機能を提供するレジスタ
        - RIP: CPUが次に実行する命令のメモリアドレスを保持するレジスタ
        - RSP: スタックの最先頭のポインタ
        - RBP: スタックの最後方のポインタ
        - RFLAGS: 命令の実行結果によって変化するフラグを集めたレジスタ
        - CR0: CPUの重要な設定を集めたレジスタ

### 3.3 はじめてのカーネル
- 本著では、ブートローダとカーネルは別個に実装する事でカーネルはブートローダの制約なく実装が可能になる。
- 以下がはじめてのカーネルである。
```c++
// extern "C"はC言語形式で関数を定義する事を意味している。
// extern "C"が必要な理由として、C++には名前修飾(マングリング)が存在するが、名前修飾を防ぐために使用する。
extern "C" void KernelMain(){
    // __asm__はアセンブリ命令を実行するイディオム
    // hlt命令はCPUを休止させる
    while (1) __asm__("hlt");
}
```
- 以下に基づいて、カーネルをコンパイルする。
```shell
$ cd $HOME/workspace/mikanos/kernel
$ git checkout osbook_day03a
$ clang++ -02 -Wall -g --target=x86_64-elf -ffreestanding -mno-red-zone -fno-exceptions -fno-rtti -std=c++17 -c main.cpp
$ ld.lld --entry KernelMain -z norelro --image-base 0x100000 --static -o kernel.elf main.o
```
- コンパイラのオプション
    - clang++
        - -02は、レベル2の最適化を行う事を表している。
        - -Wallは、警告を沢山出す。
        - -gは、デバッグ情報月でコンパイルする。
        - --target=84_64-elfは、x86_4向けの機械語を生成する。
        - -ffreestandingは、フリースタンディング環境向けのビルドを行う。
            - フリースタンディング環境向けとは、OSがない環境向けという事である。
        - -mno-red-zone
            - red-zone機能を無効にする。
        - -fno-exceptions
            - C++の例外機能を使わない。
        - -fno-rtti
            - C++の動的型情報を使わない。
        - -std=c++17
            - C++のバージョンをC++17とする。
        - -c
            - コンパイルのみ
    - lld
        - --entry KernekMain
            - KernelMain()をエントリポイントとする。
        - -z 
            - リロケーション樹夫報を読み込み専用にする機能を使わない。
        - --image-base
            - 出力されたバイナリのベースアドレスを0x100000番地とする。
        - -o kernel.elf
            - 主力ファイル名をkernel.elfとする。
        - --static
            - 静的リンクを行う。
- カーネルを読み出すブートローダは以下の通りである。
```c
EFI_FILE_PROTOCOL* kernel_file;

// カーネルファイルを開く処理
root_dir->Open(
    root_dir, 
    &kernel_file,
    L"\\kernek.elf",    // ファイルの場所を指定する。
    EFI_FILE_MODE_READ, // Readモードで開く
    0
);

UINTN file_info_size = sizeof(EFI_FILE_INFO) + sizeof(CHAR16) * 12
UINT8 file_info_buffer[file_info_size];

// ファイル情報の取得
kernel_file->GetInfo(
    kernel_file,
    &gEfiFileInfoGuid,
    &file_info_size,
    file_info_buffer;
)

// typedef struct {
//  UINT64 Size, FileSize, PhysicalSize;
//  EFI_TIME CreateTime. LastAccessTIme, ModificationTime;
//  UINT64 Attribute;
//  CHAR16 FileName[];
//} 
EFI_FILE_INFO* file_info = (EFI_FILE_INFO*)file_info_buffer;
UINTN kernel_file_size = file_info->FileSize;

EFI_PHYSICAL_ADDRES kernel_base_addr = 0x100000;

// メモリの割り当て
gBS->AllocatePages(
    AllocateAddress,    //メモリの確保の仕方
    EfiLoaderData,      //確保するメモリ領域の種別  
    (kernel_file_size + 0xfff) / 0x1000,    // 大きさ
    &kernek_base_addr   // 確保したメモリ領域のアドレスcd
);
kernel_file->Read(kernel_file, &kernel_file_size, (VOID*)kernel_base_addr);
Print(L"Kernel: 0x%0lx (%lu bytes)\n", kernel_base_addr, kernek_file_size);
```
### 3.4 ブートローダからピクセルを描く
### 3.5 カーネルからピクセルを描く
### 3.6 エラー処理をしよう
### 3.7 ポインタ入門(2): ポインタとアセンブリ言語
