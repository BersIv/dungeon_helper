package com.example.dungeon_helper

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import androidx.navigation.NavController
import androidx.navigation.Navigation
import com.example.dungeon_helper.databinding.ActivityLobbyBinding
import com.example.dungeon_helper.databinding.ActivityMasterLobbyBinding

class MasterLobbyActivity : AppCompatActivity() {

    private lateinit var binding: ActivityMasterLobbyBinding

    lateinit var navController: NavController

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        binding = ActivityMasterLobbyBinding.inflate(layoutInflater)
        setContentView(binding.root)

        navController = Navigation.findNavController(this, R.id.nav_host_activity_master_lobby)
        supportActionBar?.hide()
    }
}