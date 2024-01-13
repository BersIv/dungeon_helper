package com.example.dungeon_helper.fragments.auth

import androidx.lifecycle.ViewModel
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData

class AuthRestorePwd2ViewModel : ViewModel() {
    private val _text = MutableLiveData<String>().apply {
        value = "AuthRestorePwd2"
    }
    val text: LiveData<String> = _text
}