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
    var nickname = MutableLiveData<String>()
    val avatar = MutableLiveData<String>()

    fun getData(string: String) {
        val pairs = string.split(",")
        if (pairs.isNotEmpty()) {
            val list = ArrayList<String>()
            for (pair in pairs) {
                val pair = pair.trim().split(":")
                list.add(pair[1])
            }
            email.value = list[1]
            nickname.value = list[2]
            avatar.value = list[3]
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
