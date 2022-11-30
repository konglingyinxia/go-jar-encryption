package com.mzydz.jarencrypt;

import io.xjar.XCryptos;
import io.xjar.XEntryFilter;
import io.xjar.boot.XBoot;
import org.apache.commons.compress.archivers.jar.JarArchiveEntry;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * @author kongling
 * @package org.example
 * @date 2022/11/26 上午11:32
 * @project jar-encrypt-tools
 */
public class JarEncryptApp {
    static Logger log = LoggerFactory.getLogger(JarEncryptApp.class);

    public static void main(String[] args) {
        log.info("java进程加密开始.......");
        if (args.length != 4) {
            log.error("参数为空错误....,参数顺序为：filePath=?  outPath=?  pwd=?  startAnt=? ");
            return;
        }
        String[] filePathP = args[0].split("=");
        if (filePathP.length != 2) {
            log.error("filePath 参数错误....,参数顺序为：filePath=?  outPath=?  pwd=?  startAnt=?");
            return;
        }
        String filePath = filePathP[1];

        String[] outPathP = args[1].split("=");
        if (outPathP.length != 2) {
            log.error("outPath 参数错误....,参数顺序为：filePath=?  outPath=?  pwd=?  startAnt=? ");
            return;
        }
        String outPath = outPathP[1];

        String[] pwdP = args[2].split("=");
        if (pwdP.length != 2) {
            log.error("pwd 参数错误....,参数顺序为：filePath=?  outPath=?  pwd=?  startAnt=? ");
            return;
        }
        String pwd = pwdP[1];

        String[] startAntP = args[3].split("=");
        if (startAntP.length != 2) {
            log.error("startAnt 参数错误....,参数顺序为：filePath=?  outPath=?  pwd=?  startAnt=? ");
            return;
        }
        final String startAnt = startAntP[1];
        try {
            XBoot.encrypt(filePath, outPath, pwd, new XEntryFilter<JarArchiveEntry>() {
                @Override
                public boolean filtrate(JarArchiveEntry entry) {
                    return entry.getName().startsWith(startAnt);
                }
            });
        } catch (Exception ex) {
            log.error("java进程加密jar包失败：", ex);
        }
        log.info("java进程加密结束.......");
        System.exit(0);
    }
}
