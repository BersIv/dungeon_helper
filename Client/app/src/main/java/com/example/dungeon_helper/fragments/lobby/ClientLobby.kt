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
import com.example.dungeon_helper.databinding.FragmentClientLobbyBinding

class ClientLobby : Fragment() {

    companion object {
        fun newInstance() = ClientLobby()
    }

    private lateinit var viewModel: ClientLobbyViewModel

    private var _binding: FragmentClientLobbyBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val clientLobbyViewModel = ViewModelProvider(this)[ClientLobbyViewModel::class.java]
        _binding = FragmentClientLobbyBinding.inflate(inflater,container,false)
        val root: View = binding.root
        val textView: TextView = binding.textPlay
        clientLobbyViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val viewChar1 = binding.viewChar1
        viewChar1.setOnClickListener{
            (activity as ClientLobbyActivity).navController.navigate(R.id.action_clientLobby_to_clientLobbyViewChar)
        }
        val viewChar2 = binding.viewChar2
        val backBtn = binding.backBtn
        val infoBtn = binding.infoBtn
        infoBtn.setOnClickListener {
            (activity as ClientLobbyActivity).navController.navigate(R.id.action_clientLobby_to_clientLobbyInfoAllChapters)
        }
        val viewBtn = binding.viewBtn
        viewBtn.setOnClickListener {
            (activity as ClientLobbyActivity).navController.navigate(R.id.action_clientLobby_to_clientLobbyViewChar)
        }
        val upLvlBtn = binding.upLvlBtn
    }


}