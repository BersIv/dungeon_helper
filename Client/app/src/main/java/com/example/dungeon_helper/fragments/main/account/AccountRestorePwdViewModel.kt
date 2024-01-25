package com.example.dungeon_helper.fragments.main.account

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class AccountRestorePwdViewModel : ViewModel() {
    private val _text =  MutableLiveData<String>().apply {
        value = "AccountRestorePwd"
    }
    val text: LiveData<String> = _text
}