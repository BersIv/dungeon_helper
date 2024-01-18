package com.example.dungeon_helper

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import androidx.navigation.NavController
import androidx.navigation.Navigation
import com.example.dungeon_helper.databinding.ActivityClientLobbyBinding
import com.example.dungeon_helper.databinding.ActivityLobbyBinding

class ClientLobbyActivity : AppCompatActivity() {

    private lateinit var binding: ActivityClientLobbyBinding

    lateinit var navController: NavController

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        binding = ActivityClientLobbyBinding.inflate(layoutInflater)
        setContentView(binding.root)

        navController = Navigation.findNavController(this, R.id.nav_host_activity_client_lobby)
        supportActionBar?.hide()
    }

}