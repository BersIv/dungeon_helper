package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MasterLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentMasterLobbyChooseBinding

class MasterLobbyChoose : Fragment() {

    companion object {
        fun newInstance() = MasterLobbyChoose()
    }

    private lateinit var viewModel: MasterLobbyChooseViewModel
    private var _binding: FragmentMasterLobbyChooseBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val masterLobbyChooseViewModel = ViewModelProvider(this)[MasterLobbyChooseViewModel::class.java]
        _binding = FragmentMasterLobbyChooseBinding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textPlay
        masterLobbyChooseViewModel.text.observe(viewLifecycleOwner){
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
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyChoose_to_masterLobbyActChoose)
        }

        val viewChar2 = binding.viewChar2
        val viewChar3 = binding.viewChar3
        val backBtn = binding.backBtn
        backBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyChoose_to_masterLobby)
        }
    }

}