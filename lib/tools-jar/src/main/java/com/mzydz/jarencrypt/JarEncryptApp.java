package com.mzydz.jarencrypt;

import io.xjar.XCryptos;
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
        if (args.length != 3) {
            log.error("参数为空错误....,参数顺序为：filePath=? pwd=?  outPath=? ");
            return;
        }
        String[] filePathP = args[0].split("=");
        if (filePathP.length != 2) {
            log.error("filePath 参数错误....,参数顺序为：filePath=? pwd=?  outPath=? ");
            return;
        }
        String filePath = filePathP[1];
        String[] pwdP = args[1].split("=");
        if (pwdP.length != 2) {
            log.error("pwd 参数错误....,参数顺序为：filePath=? pwd=?  outPath=? ");
            return;
        }
        String pwd = pwdP[1];
        String[] outPathP = args[2].split("=");
        if (outPathP.length != 2) {
            log.error("outPath 参数错误....,参数顺序为：filePath=? pwd=?  outPath=? ");
            return;
        }
        log.info("java进程加密开始.......");
        String outPath = outPathP[1];
        try {
            XCryptos.encryption()
                    .from(filePath)
                    .use(pwd)
                    .exclude("/static/**/*")
                    .exclude("/excel/**/*")
                    .exclude("/templates/**/*")
                    .exclude("/META-INF/resources/**/*")
                    .to(outPath);
        } catch (Exception ex) {
            log.error("java加密jar包失败：", ex);
        }
        log.info("java进程加密结束.......");
    }
}
