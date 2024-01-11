package com.example.dungeon_helper.fragments.auth

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class AuthRestorePwdViewModel : ViewModel() {

    private val _text = MutableLiveData<String>().apply {
        value = "AuthRestorePwd"
    }
    val text: LiveData<String> = _text
}