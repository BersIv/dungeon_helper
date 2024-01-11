package com.example.dungeon_helper.fragments.auth

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class AuthMainViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "Auth"
    }
    val text: LiveData<String> = _text
}