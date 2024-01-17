package com.example.dungeon_helper

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.navigation.NavController
import androidx.navigation.Navigation
import com.example.dungeon_helper.databinding.ActivityLobbyBinding

class LobbyActivity: AppCompatActivity() {

    private lateinit var binding: ActivityLobbyBinding

    lateinit var navController: NavController

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        binding = ActivityLobbyBinding.inflate(layoutInflater)
        setContentView(binding.root)

        navController = Navigation.findNavController(this, R.id.nav_host_activity_lobby)
    }
}