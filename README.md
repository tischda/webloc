webloc decompiler in GO
=======================

Requirements
------------
* GO 1.4.2



Issues
------
* Does not read resource forks, only data in plist file


References
----------
* https://github.com/DHowett/go-plist
* https://github.com/ggilder/weirdfs


Decompile binary weblocs
========================
hansolo:~ daniel$ hexdump -C binary-plist.webloc
00000000  62 70 6c 69 73 74 30 30  d1 01 02 53 55 52 4c 5f  |bplist00...SURL_|
00000010  10 1a 68 74 74 70 3a 2f  2f 77 77 77 2e 6b 65 6b  |..http://www.kek|
00000020  61 6f 73 78 2e 63 6f 6d  2f 65 6e 2f 08 0b 0f 00  |aosx.com/en/....|
00000030  00 00 00 00 00 01 01 00  00 00 00 00 00 00 03 00  |................|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 2c     |..............,|
0000004f
hansolo:~

hansolo:~ daniel$ derez binary-plist.webloc
data 'drag' (128) {
    $"0000 0001 0000 0000 0000 0000 0000 0003"            /* ................ */
    $"5445 5854 0000 0100 0000 0000 0000 0000"            /* TEXT............ */
    $"7572 6C20 0000 0100 0000 0000 0000 0000"            /* url ............ */
    $"7572 6C6E 0000 0100 0000 0000 0000 0000"            /* urln............ */
};

data 'url ' (256) {
    $"6874 7470 3A2F 2F77 7777 2E6B 656B 616F"            /* http://www.kekao */
    $"7378 2E63 6F6D 2F65 6E2F"                           /* sx.com/en/ */
};

data 'TEXT' (256) {
    $"6874 7470 3A2F 2F77 7777 2E6B 656B 616F"            /* http://www.kekao */
    $"7378 2E63 6F6D 2F65 6E2F"                           /* sx.com/en/ */
};

data 'urln' (256) {
    $"4B65 6B61 202D 2074 6865 2066 7265 6520"            /* Keka - the free  */
    $"4D61 6320 4F53 2058 2066 696C 6520 6172"            /* Mac OS X file ar */
    $"6368 6976 6572"                                     /* chiver */
};


hansolo:webloc daniel$ DeRez -e -only 'urln' binary-plist.webloc
data 'urln' (256) {
    $"4B65 6B61 202D 2074 6865 2066 7265 6520"            /* Keka - the free  */
    $"4D61 6320 4F53 2058 2066 696C 6520 6172"            /* Mac OS X file ar */
    $"6368 6976 6572"                                     /* chiver */
};
