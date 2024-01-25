package com.example.dungeon_helper.shared

import android.content.Intent
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.example.dungeon_helper.MainActivity
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.WebSocket
import okhttp3.WebSocketListener
import java.io.BufferedReader
import java.io.IOException
import java.io.InputStream
import java.io.InputStreamReader
import java.io.Reader
import java.io.StringWriter
import java.io.Writer

class SharedViewModel: ViewModel() {

    val newCharacter = MutableLiveData<Character>()
    val getCharacter = MutableLiveData<GetCharacter>()
    val characterClasses = MutableLiveData<List<CharacterClass>>()
    val races = MutableLiveData<List<Race>>()
    val subraces = MutableLiveData<List<Subrace>>()
    val alignments = MutableLiveData<List<Alignment>>()
    val skills = MutableLiveData<List<CharacterSkill>>()
    val allChar = MutableLiveData<List<AllCharacter>>()

    val token = MutableLiveData<String>()
    val email = MutableLiveData<String>()
    val nickname = MutableLiveData<String>()
    val avatar = MutableLiveData<String>()
    val mailChangePwd = MutableLiveData<String>()
    val mailRestorePwd = MutableLiveData<String>()

    private val _webSocketOpen = MutableLiveData<Boolean>()
    val socket = MutableLiveData<WebSocket>()
    val idToChange = MutableLiveData<String>()

    val webSocketOpen: LiveData<Boolean> get() = _webSocketOpen

    fun setWebSocketOpen(open: Boolean) {
        _webSocketOpen.postValue(open)
    }

    init {
        

        // Устанавливаем начальные значения для newCharacter
        newCharacter.value = Character(
            hp = 100,
            exp = 0,
            avatar = "",
            charName = "",
            sex = true,  // Пример начального значения для boolean
            weight = 0,
            height = 0,
            charClass = CharacterClass(0, ""),
            race = Race(0, ""),
            subrace = Subrace(0, "", Stats(0, 0, 0, 0, 0, 0)),
            stats = Stats(0, 0, 0, 0, 0, 0),
            addLanguage = "",
            characterSkills = listOf(CharacterSkill(2, "Акробатика")),
            alignment = Alignment(1, ""),
            ideals = "",
            weaknesses = "",
            traits = "",
            allies = "",
            organizations = "",
            enemies = "",
            story = "",
            goals = "",
            treasures = "",
            notes = ""
        )


        getCharacter.value = GetCharacter(
            id = 0,
            hp = 100,
            lvl = 1,
            exp = 0,
            avatar = "",
            charName = "",
            sex = true,  // Пример начального значения для boolean
            weight = 0,
            height = 0,
            charClass = "",
            race = "",
            subrace = "",
            stats = Stats(0, 0, 0, 0, 0, 0),
            addLanguage = "",
            characterSkills = "",
            alignment = "",
            ideals = "",
            weaknesses = "",
            traits = "",
            allies = "",
            organizations = "",
            enemies = "",
            story = "",
            goals = "",
            treasures = "",
            notes = ""
        )
    }

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

    var webSocket: WebSocket? = null



    private fun closeWebSocket() {
        // Close WebSocket connection
        webSocket?.close(1000, "Closed by user")
    }
}

