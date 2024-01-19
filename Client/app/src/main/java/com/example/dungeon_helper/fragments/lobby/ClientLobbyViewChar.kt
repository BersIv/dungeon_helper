package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.ClientLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentClientLobbyViewCharBinding

class ClientLobbyViewChar : Fragment() {

    companion object {
        fun newInstance() = ClientLobbyViewChar()
    }

    private lateinit var viewModel: ClientLobbyViewCharViewModel
    private var _binding: FragmentClientLobbyViewCharBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val clientLobbyViewCharViewModel = ViewModelProvider(this)[ClientLobbyViewCharViewModel::class.java]
        _binding = FragmentClientLobbyViewCharBinding.inflate(inflater, container, false)
        val root: View = binding.root
        return root
    }
    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val backBtn = binding.backBtn
        backBtn.setOnClickListener {
            (activity as ClientLobbyActivity).navController.navigate(R.id.action_clientLobbyViewChar_to_clientLobby)
        }
    }

}