package com.example.dungeon_helper

import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import java.io.BufferedReader
import java.io.IOException
import java.io.InputStream
import java.io.InputStreamReader
import java.io.Reader
import java.io.StringWriter
import java.io.Writer

class SharedViewModel: ViewModel() {

    val token = MutableLiveData<String>()
    val email = MutableLiveData<String>()
    val nickname = MutableLiveData<String>()
    val avatar = MutableLiveData<String>()
    val mailChangePwd = MutableLiveData<String>()
    val mailRestorePwd = MutableLiveData<String>()

    fun getData(string: String) {
        println(string)
        val pairs = string.split(",")
        println(pairs)
        if (pairs.isNotEmpty()) {
            val list = ArrayList<String>()
            for (pair in pairs) {
                val pair = pair.trim().split(":")
                println(pair)
                list.add(pair[1])
            }
            println(list)
            println(list[1])
            email.value = list[1].removeSurrounding("\"")
            nickname.value = list[2].removeSurrounding("\"")
            avatar.value = list[3].removeSurrounding("\"")
        }

    }

    @Throws(IOException::class)
    fun convertStreamToString(inputStream: InputStream?): String {
        return if (inputStream != null) {
            val writer: Writer = StringWriter()
            val buffer = CharArray(1024)
            try {
                val reader: Reader = BufferedReader(InputStreamReader(inputStream, "UTF-8"), 1024)
                var n: Int
                while (reader.read(buffer).also { n = it } != -1) {
                    writer.write(buffer, 0, n)
                }
            } finally {
                inputStream.close()
            }
            writer.toString()
        } else {
            ""
        }
    }

    fun getToken(): String? {
        return token.value
    }

    fun getEmail(): String? {
        return email.value
    }

    fun getNick(): String? {
        return nickname.value
    }

    fun getAvatar(): String? {
        return avatar.value
    }


}
