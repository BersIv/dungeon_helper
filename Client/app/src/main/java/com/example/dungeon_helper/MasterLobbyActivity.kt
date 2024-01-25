package com.example.dungeon_helper

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.navigation.NavController
import androidx.navigation.Navigation
import com.example.dungeon_helper.databinding.ActivityLobbyBinding
import com.example.dungeon_helper.databinding.ActivityMasterLobbyBinding
import com.example.dungeon_helper.shared.SharedViewModel
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.Response
import okhttp3.WebSocket
import okhttp3.WebSocketListener
import okio.ByteString
import okio.ByteString.Companion.decodeHex
import java.util.concurrent.TimeUnit

class MasterLobbyActivity : AppCompatActivity() {

    private lateinit var binding: ActivityMasterLobbyBinding

    lateinit var navController: NavController

    private val socketUrl = "ws://194.247.187.44:5000/lobby/create?lobbyName=1&lobbyPassword=1&amount=2"
    private var webSocket: WebSocket? = null

    lateinit var sharedViewModel: SharedViewModel

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        sharedViewModel = ViewModelProvider(this)[SharedViewModel::class.java]

        binding = ActivityMasterLobbyBinding.inflate(layoutInflater)
        setContentView(binding.root)

        navController = Navigation.findNavController(this, R.id.nav_host_activity_master_lobby)
        supportActionBar?.hide()

        initWebSocket()

//        val client = OkHttpClient.Builder()
//            .readTimeout(3, TimeUnit.SECONDS)
//            //.sslSocketFactory() - ? нужно ли его указывать дополнительно
//            .build()
//        val request = Request.Builder()
//            .url(socketUrl) // 'wss' - для защищенного канала
//            .build()
//        val wsListener = EchoWebSocketListener ()
//        val webSocket = client.newWebSocket(request, wsListener) // this provide to make 'Open ws connection'
    }

    private fun initWebSocket() {
        val request = Request.Builder()
            .url(socketUrl)
            .build()

        val webSocketListener = object : WebSocketListener() {
            override fun onOpen(webSocket: WebSocket, response: okhttp3.Response) {

                sharedViewModel.setWebSocketOpen(true)
                println("change_exp id{} exp{}")
            }

            override fun onMessage(webSocket: WebSocket, text: String) {
                println(text)
            }

            override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {

                sharedViewModel.setWebSocketOpen(false)

                runOnUiThread {
                    val intent = Intent(this@MasterLobbyActivity, MainActivity::class.java)
                    startActivity(intent)
                }
            }

            override fun onFailure(webSocket: WebSocket, t: Throwable, response: okhttp3.Response?) {

                sharedViewModel.setWebSocketOpen(false)

                if (response != null) {
                    println("${response.code} ${response.message}")
                }

                runOnUiThread {
                    val intent = Intent(this@MasterLobbyActivity, MainActivity::class.java)
                    startActivity(intent)
                }
            }
        }
        var token = ""
        token = intent.getStringExtra("token").toString()
        val client = OkHttpClient.Builder().addInterceptor { chain ->
            val original = chain.request()
            val authorized = original.newBuilder()
                .header("Authorization", token)
                .build()
            chain.proceed(authorized)
        }
            .build()
        // Create WebSocket connection
        webSocket = client.newWebSocket(request, webSocketListener)
        sharedViewModel.socket.value = webSocket
    }

    private fun closeWebSocket() {
        // Close WebSocket connection
        webSocket?.close(1000, "Closed by user")
    }
}

//private class EchoWebSocketListener : WebSocketListener() {
//    override fun onOpen(webSocket: WebSocket, response: Response) {
//        webSocket.send("Hello, it's SSaurel !")
//        webSocket.send("What's up ?")
//        webSocket.send("deadbeef".decodeHex())
//        webSocket.close(NORMAL_CLOSURE_STATUS, "Goodbye !")
//    }
//
//    fun onMessage(webSocket: WebSocket?, text: String?) {
//        output("Receiving : " + text!!)
//    }
//
//    fun onMessage(webSocket: WebSocket?, bytes: ByteString?) {
//        output("Receiving bytes : " + bytes!!.hex())
//    }
//
//    fun onClosing(webSocket: WebSocket?, code: Int, reason: String?) {
//        webSocket!!.close(NORMAL_CLOSURE_STATUS, null)
//        output("Closing : $code / $reason")
//    }
//
//    fun onFailure(webSocket: WebSocket, t: Throwable, response: Response) {
//        output("Error : " + t.message)
//    }
//
//    companion object {
//        private val NORMAL_CLOSURE_STATUS = 1000
//    }
//
//    private fun output(txt: String) {
//        Log.v("WSS", txt)
//    }
//}