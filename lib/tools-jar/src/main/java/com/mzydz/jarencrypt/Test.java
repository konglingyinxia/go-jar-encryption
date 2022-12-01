package com.mzydz.jarencrypt;

import io.xjar.XEntryFilter;
import io.xjar.boot.XBoot;
import io.xjar.filter.XAntEntryFilter;
import io.xjar.filter.XAnyEntryFilter;
import org.apache.commons.compress.archivers.jar.JarArchiveEntry;

import java.io.File;
import java.util.ArrayList;
import java.util.List;

/**
 * @author kongling
 * @package com.mzydz.jarencrypt
 * @date 2022/11/30 下午7:46
 * @project jar-encrypt-tools
 */
public class Test {
    public static void main(String[] args) throws Exception {
        String password = "io.xjar";
        File plaintext = new File("/home/kongling/桌面/test/project-platform-item-1.0-SNAPSHOT.jar");
        File encrypted = new File("/home/kongling/桌面/test/ad_publish_bank_test/out/project-platform-item-1.0-SNAPSHOT.jar");

        XBoot.encrypt(plaintext, encrypted, password, new XEntryFilter<JarArchiveEntry>() {
            @Override
            public boolean filtrate(JarArchiveEntry entry) {
                System.out.println(entry.getName());
                return entry.getName().startsWith("com/mzydz/");
            }
        });
    }
}
