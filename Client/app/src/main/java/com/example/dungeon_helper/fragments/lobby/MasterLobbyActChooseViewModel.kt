package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class MasterLobbyActChooseViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "PLAY_MASTER_ACT"
    }
    val text: LiveData<String> = _text
}