package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.dungeon_helper.MasterLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentMasterLobbyViewCharBinding

class MasterLobbyViewChar : Fragment() {

    companion object {
        fun newInstance() = MasterLobbyViewChar()
    }

    private lateinit var viewModel: MasterLobbyViewCharViewModel
    private var _binding: FragmentMasterLobbyViewCharBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val masterLobbyViewCharViewModel = ViewModelProvider(this)[MasterLobbyViewModel::class.java]
        _binding = FragmentMasterLobbyViewCharBinding.inflate(inflater, container, false)
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
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyViewChar_to_masterLobby)
        }
    }



}